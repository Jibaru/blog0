'use client';

import { Bookmark, Calendar, Heart, LogIn, MessageCircle, Share, User, Plus, Play, Pause } from 'lucide-react';
import Link from 'next/link';
import { useCallback, useEffect, useRef, useState } from 'react';
import LoginModal from '@/components/auth/LoginModal';
import UserMenu from '@/components/auth/UserMenu';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { type PostItem, ApiError } from '@/lib/api-client';
import { useAuthStore } from '@/store/authStore';
import { useToast } from '@/components/ui/toast';
import FollowButton from '@/components/FollowButton';

export default function Home() {
  const [posts, setPosts] = useState<PostItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [loadingMore, setLoadingMore] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [showLoginModal, setShowLoginModal] = useState(false);
  const [currentPostIndex, setCurrentPostIndex] = useState(0);
  const [currentPage, setCurrentPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const loadingRef = useRef(false);
  
  // TikTok-style navigation refs
  const containerRef = useRef<HTMLDivElement>(null);
  const touchStartY = useRef(0);
  const touchEndY = useRef(0);
  const isTransitioning = useRef(false);
  
  // Audio player state
  const audioRef = useRef<HTMLAudioElement | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentTime, setCurrentTime] = useState(0);
  const [duration, setDuration] = useState(0);
  const [currentAudioUrl, setCurrentAudioUrl] = useState<string | null>(null);
  const [hasUserInteracted, setHasUserInteracted] = useState(false);
  const [showPlayPrompt, setShowPlayPrompt] = useState(false);
  
  // Track local like counts for posts
  const [postLikeCounts, setPostLikeCounts] = useState<Record<string, number>>({});

  const { 
    isAuthenticated, 
    getApiClient, 
    isPostLiked, 
    isPostBookmarked, 
    updatePostLike, 
    updatePostBookmark 
  } = useAuthStore();
  const { showToast } = useToast();

  const fetchPosts = useCallback(async (page: number, append = false) => {
    try {
      if (page === 1) {
        setLoading(true);
      } else {
        setLoadingMore(true);
      }

      const apiClient = getApiClient();
      const response = await apiClient.listPosts({ page, per_page: 5 });
      
      if (append) {
        setPosts(prev => [...prev, ...response.items]);
      } else {
        setPosts(response.items);
      }

      // Initialize like counts for new posts
      const newLikeCounts: Record<string, number> = {};
      response.items.forEach(post => {
        newLikeCounts[post.slug] = post.like_count;
      });
      
      if (append) {
        setPostLikeCounts(prev => ({ ...prev, ...newLikeCounts }));
      } else {
        setPostLikeCounts(newLikeCounts);
      }

      // Check if there are more posts to load
      const totalPages = Math.ceil(response.total / 5);
      setHasMore(page < totalPages);
      
    } catch (err) {
      const errorMessage = err instanceof ApiError ? err.message : 'Failed to fetch posts';
      setError(errorMessage);
      showToast(errorMessage);
    } finally {
      setLoading(false);
      setLoadingMore(false);
    }
  }, [getApiClient, showToast]);

  // Audio player functions
  const playAudio = useCallback(async (audioUrl: string, forceAttempt = false) => {
    if (!audioRef.current) {
      audioRef.current = new Audio();
      
      // Set up audio event listeners
      audioRef.current.addEventListener('loadedmetadata', () => {
        setDuration(audioRef.current?.duration || 0);
      });
      
      audioRef.current.addEventListener('timeupdate', () => {
        setCurrentTime(audioRef.current?.currentTime || 0);
      });
      
      audioRef.current.addEventListener('ended', () => {
        setIsPlaying(false);
        setCurrentTime(0);
      });
      
      audioRef.current.addEventListener('play', () => {
        setIsPlaying(true);
      });
      
      audioRef.current.addEventListener('pause', () => {
        setIsPlaying(false);
      });
    }
    
    // If same audio, just play/resume
    if (currentAudioUrl === audioUrl) {
      try {
        await audioRef.current.play();
      } catch (error) {
        if (!hasUserInteracted && !forceAttempt) {
          setShowPlayPrompt(true);
        }
        console.warn('Audio play failed:', error);
      }
      return;
    }
    
    // New audio - load and attempt to play
    setCurrentAudioUrl(audioUrl);
    audioRef.current.src = audioUrl;
    audioRef.current.currentTime = 0;
    setCurrentTime(0);
    
    try {
      await audioRef.current.play();
    } catch (error) {
      if (!hasUserInteracted && !forceAttempt) {
        setShowPlayPrompt(true);
      } else {
        console.warn('Audio play failed:', error);
      }
    }
  }, [currentAudioUrl, hasUserInteracted]);

  const pauseAudio = useCallback(() => {
    if (audioRef.current && !audioRef.current.paused) {
      audioRef.current.pause();
    }
  }, []);

  const toggleAudio = useCallback(async () => {
    if (!audioRef.current) return;
    
    // Mark user interaction
    if (!hasUserInteracted) {
      setHasUserInteracted(true);
      setShowPlayPrompt(false);
    }
    
    if (isPlaying) {
      pauseAudio();
    } else {
      if (currentAudioUrl) {
        await playAudio(currentAudioUrl, true);
      } else if (posts[currentPostIndex]?.summary_audio_url) {
        await playAudio(posts[currentPostIndex].summary_audio_url, true);
      }
    }
  }, [isPlaying, pauseAudio, hasUserInteracted, currentAudioUrl, posts, currentPostIndex, playAudio]);

  const handleUserInteraction = useCallback(() => {
    if (!hasUserInteracted) {
      setHasUserInteracted(true);
      setShowPlayPrompt(false);
    }
  }, [hasUserInteracted]);

  const seekTo = useCallback((time: number) => {
    if (audioRef.current) {
      audioRef.current.currentTime = time;
      setCurrentTime(time);
    }
  }, []);

  // TikTok-style navigation functions
  const goToNextPost = useCallback(() => {
    if (isTransitioning.current) return;
    
    // If we're at the last post and have more content, load next page
    if (currentPostIndex === posts.length - 1 && hasMore && !loadingMore) {
      loadingRef.current = true;
      const nextPage = currentPage + 1;
      setCurrentPage(nextPage);
      fetchPosts(nextPage, true).catch((error) => {
        if (error instanceof ApiError) {
          showToast(error.message);
        } else {
          showToast('Failed to load more posts');
        }
      }).finally(() => {
        loadingRef.current = false;
      });
    }
    
    // Navigate to next post if available
    if (currentPostIndex < posts.length - 1) {
      isTransitioning.current = true;
      setCurrentPostIndex(prev => prev + 1);
      setTimeout(() => {
        isTransitioning.current = false;
      }, 200);
    }
  }, [currentPostIndex, posts.length, hasMore, loadingMore, currentPage, fetchPosts, showToast]);

  // Handle audio playback when current post changes
  useEffect(() => {
    if (posts.length === 0) return;
    
    const currentPost = posts[currentPostIndex];
    if (currentPost?.summary_audio_url) {
      // Post has audio - play it automatically
      playAudio(currentPost.summary_audio_url);
    } else {
      // Post has no audio - pause current audio
      pauseAudio();
      setCurrentAudioUrl(null);
    }
  }, [currentPostIndex, posts, playAudio, pauseAudio]);

  // Cleanup audio on unmount
  useEffect(() => {
    return () => {
      if (audioRef.current) {
        audioRef.current.pause();
        audioRef.current = null;
      }
    };
  }, []);

  const goToPreviousPost = useCallback(() => {
    if (isTransitioning.current) return;
    
    if (currentPostIndex > 0) {
      isTransitioning.current = true;
      setCurrentPostIndex(prev => prev - 1);
      setTimeout(() => {
        isTransitioning.current = false;
      }, 200);
    }
  }, [currentPostIndex]);

  // Touch gesture handling
  const handleTouchStart = useCallback((e: React.TouchEvent) => {
    touchStartY.current = e.touches[0].clientY;
    handleUserInteraction();
  }, [handleUserInteraction]);

  const handleTouchEnd = useCallback((e: React.TouchEvent) => {
    touchEndY.current = e.changedTouches[0].clientY;
    const deltaY = touchStartY.current - touchEndY.current;
    const minSwipeDistance = 50;

    if (Math.abs(deltaY) > minSwipeDistance) {
      if (deltaY > 0) {
        // Swiped up - next post
        goToNextPost();
      } else {
        // Swiped down - previous post
        goToPreviousPost();
      }
    }
  }, [goToNextPost, goToPreviousPost]);

  // Wheel/scroll handling for desktop
  const handleWheel = useCallback((e: WheelEvent) => {
    e.preventDefault();
    handleUserInteraction();
    
    if (Math.abs(e.deltaY) > 10) {
      if (e.deltaY > 0) {
        // Scrolled down - next post
        goToNextPost();
      } else {
        // Scrolled up - previous post
        goToPreviousPost();
      }
    }
  }, [goToNextPost, goToPreviousPost, handleUserInteraction]);

  // Keyboard navigation
  const handleKeyDown = useCallback((e: KeyboardEvent) => {
    handleUserInteraction();
    if (e.key === 'ArrowDown' || e.key === ' ') {
      e.preventDefault();
      goToNextPost();
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      goToPreviousPost();
    }
  }, [goToNextPost, goToPreviousPost, handleUserInteraction]);

  const handleLike = async (slug: string) => {
    handleUserInteraction();
    
    if (!isAuthenticated) {
      setShowLoginModal(true);
      return;
    }

    try {
      const apiClient = getApiClient();
      const response = await apiClient.toggleLike(slug);
      
      // Update profile state
      updatePostLike(slug, response.liked);
      
      // Update like count
      setPostLikeCounts(prev => ({
        ...prev,
        [slug]: response.likes_count
      }));
      
    } catch (err) {
      console.error('Failed to toggle like:', err);
      if (err instanceof ApiError) {
        showToast(err.message);
      } else {
        showToast('Failed to update like');
      }
    }
  };

  const handleBookmark = async (slug: string) => {
    handleUserInteraction();
    
    if (!isAuthenticated) {
      setShowLoginModal(true);
      return;
    }

    try {
      const apiClient = getApiClient();
      const isCurrentlyBookmarked = isPostBookmarked(slug);
      
      if (isCurrentlyBookmarked) {
        await apiClient.unbookmarkPost(slug);
        updatePostBookmark(slug, false);
      } else {
        await apiClient.bookmarkPost(slug);
        updatePostBookmark(slug, true);
      }
      
    } catch (err) {
      console.error('Failed to toggle bookmark:', err);
      if (err instanceof ApiError) {
        showToast(err.message);
      } else {
        showToast('Failed to update bookmark');
      }
    }
  };

  useEffect(() => {
    fetchPosts(1);
  }, [fetchPosts]);

  // Set up event listeners for navigation
  useEffect(() => {
    const container = containerRef.current;
    if (!container) return;

    // Add wheel event listener
    container.addEventListener('wheel', handleWheel, { passive: false });
    
    // Add keyboard event listener
    document.addEventListener('keydown', handleKeyDown);

    // Cleanup
    return () => {
      container.removeEventListener('wheel', handleWheel);
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [handleWheel, handleKeyDown]);

  // Auto-load next page when approaching end
  useEffect(() => {
    if (currentPostIndex >= posts.length - 2 && hasMore && !loadingMore && !loadingRef.current) {
      loadingRef.current = true;
      const nextPage = currentPage + 1;
      setCurrentPage(nextPage);
      fetchPosts(nextPage, true).catch((error) => {
        if (error instanceof ApiError) {
          showToast(error.message);
        } else {
          showToast('Failed to load more posts');
        }
      }).finally(() => {
        loadingRef.current = false;
      });
    }
  }, [currentPostIndex, posts.length, hasMore, loadingMore, currentPage, fetchPosts, showToast]);

  if (loading) {
    return (
      <div className="post-container flex items-center justify-center mesh-background">
        <div className="text-2xl accent-text font-bold">Loading posts...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="post-container flex items-center justify-center mesh-background">
        <div className="text-center space-y-4">
          <div className="text-xl text-[#FE2C55]">Error: {error}</div>
          <Button
            onClick={() => window.location.reload()}
            className="bg-[#FE2C55] text-white border-0 hover:bg-[#FE2C55]/80"
          >
            Retry
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="bg-black mesh-background">
      {/* Slim top bar */}
      <header className="fixed top-0 left-0 right-0 z-50 h-16 flex items-center justify-between px-4 border-b border-transparent bg-black/80 backdrop-blur-sm">
        <h1 className="accent-text text-xl font-bold">Blog0</h1>

        {/* Auth section */}
        <div className="flex items-center space-x-3">
          {isAuthenticated && (
            <Link href="/create">
              <Button
                variant="ghost"
                size="sm"
                className="text-white hover:text-[#25F4EE] hover:bg-white/10"
              >
                <Plus className="h-4 w-4 mr-2" />
                Create
              </Button>
            </Link>
          )}
          
          {isAuthenticated ? (
            <UserMenu />
          ) : (
            <Button
              variant="ghost"
              size="sm"
              className="text-white hover:text-[#FE2C55] hover:bg-white/10"
              onClick={() => setShowLoginModal(true)}
            >
              <LogIn className="h-4 w-4 mr-2" />
              Sign In
            </Button>
          )}
        </div>
      </header>

      {/* TikTok-style single post container */}
      <div 
        ref={containerRef}
        className="tiktok-container"
        onTouchStart={handleTouchStart}
        onTouchEnd={handleTouchEnd}
      >
        {posts.length === 0 ? (
          <div className="post-item">
            <div className="text-center space-y-4">
              <div className="text-2xl text-[#AFAFAF]">No posts available</div>
              <div className="accent-text text-lg">Check back later!</div>
            </div>
          </div>
        ) : (
          <>
            {/* Render current post and adjacent posts for smooth transitions */}
            {posts.map((post, index) => {
              const position = index - currentPostIndex;
              let positionClass = '';
              
              if (position === 0) {
                positionClass = 'current-post';
              } else if (position === 1) {
                positionClass = 'next-post';
              } else if (position === -1) {
                positionClass = 'prev-post';
              } else {
                // Hide posts that are not adjacent
                return null;
              }
              
              return (
              <div 
                key={post.slug}
                className={`post-item-tiktok ${positionClass}`}
                style={{
                  transform: `translateY(${position * 100}%)`,
                }}
              >
                {/* Background content area */}
                <div className="absolute inset-0 bg-gradient-to-b from-black/0 via-black/20 to-black/80" />

                {/* Main content */}
                <div className="relative z-10 w-full h-full flex flex-col">
                  {/* Content container - centered vertically */}
                  <div className="flex-1 flex flex-col justify-center px-4 pr-20 max-w-md mx-auto w-full">
                    {/* Author info */}
                    <div className="mb-8 flex items-center space-x-3">
                      <Avatar className="w-12 h-12 border-2 border-white/20">
                        <AvatarFallback className="bg-[#FE2C55] text-white font-bold">
                          {post.author.charAt(0).toUpperCase()}
                        </AvatarFallback>
                      </Avatar>
                      <div>
                        <div className="text-white font-semibold">{post.author}</div>
                        <div className="text-[#AFAFAF] text-sm flex items-center">
                          <User className="h-3 w-3 mr-1" />@{post.author}
                        </div>
                      </div>
                    </div>

                    {/* Post info */}
                    <Link href={`/post/${post.slug}`} className="block">
                      <div className="space-y-4">
                        {/* Post title */}
                        <h2 className="text-2xl md:text-3xl font-bold leading-tight text-white break-words">
                          {post.title}
                        </h2>

                        {/* Post summary */}
                        {post.summary && (
                          <p className="text-[#AFAFAF] text-base leading-relaxed line-clamp-4">
                            {post.summary}
                          </p>
                        )}

                        {/* Post metadata */}
                        <div className="flex items-center space-x-4 text-[#AFAFAF]">
                          {post.published_at && (
                            <div className="flex items-center caption-small">
                              <Calendar className="h-3 w-3 mr-1" />
                              {new Date(post.published_at).toLocaleDateString()}
                            </div>
                          )}
                        </div>
                        
                        {/* Tags */}
                        {post.tags && post.tags.length > 0 && (
                          <div className="flex flex-wrap gap-2">
                            {post.tags.slice(0, 3).map((tag, tagIndex) => (
                              <span key={tagIndex} className="text-sm text-[#FE2C55] font-medium">
                                #{tag}
                              </span>
                            ))}
                            {post.tags.length > 3 && (
                              <span className="text-sm text-[#AFAFAF]">+{post.tags.length - 3} more</span>
                            )}
                          </div>
                        )}

                        {/* Engagement preview */}
                        <div className="flex items-center space-x-6 text-[#AFAFAF] text-sm">
                          <span className="flex items-center">
                            <Heart className="h-4 w-4 mr-2" />
                            {postLikeCounts[post.slug] ?? post.like_count} likes
                          </span>
                          <span className="flex items-center">
                            <MessageCircle className="h-4 w-4 mr-2" />
                            {post.comment_count} comments
                          </span>
                        </div>
                      </div>
                    </Link>
                  </div>
                </div>

                {/* Right sidebar with action buttons */}
                <div className="absolute right-3 top-1/2 -translate-y-1/2 flex flex-col space-y-4 z-20">
                  {/* Like button */}
                  <div className="flex flex-col items-center space-y-2">
                    <Button
                      size="lg"
                      variant="ghost"
                      className={`w-12 h-12 rounded-full p-0 transition-smooth hover:scale-110 ${
                        isPostLiked(post.slug)
                          ? 'text-[#FE2C55] glow-accent'
                          : 'text-white hover:text-[#FE2C55] hover:bg-white/10'
                      }`}
                      onClick={() => handleLike(post.slug)}
                    >
                      <Heart className={`h-7 w-7 ${isPostLiked(post.slug) ? 'fill-current' : ''}`} />
                    </Button>
                    <span className="caption-small text-white font-medium">
                      {postLikeCounts[post.slug] ?? post.like_count}
                    </span>
                  </div>

                  {/* Comment button */}
                  <div className="flex flex-col items-center space-y-2">
                    <Link href={`/post/${post.slug}`}>
                      <Button
                        size="lg"
                        variant="ghost"
                        className="w-12 h-12 rounded-full p-0 text-white hover:text-[#25F4EE] hover:bg-white/10 transition-smooth hover:scale-110"
                      >
                        <MessageCircle className="h-7 w-7" />
                      </Button>
                    </Link>
                    <span className="caption-small text-white font-medium">
                      {post.comment_count}
                    </span>
                  </div>

                  {/* Bookmark button */}
                  <div className="flex flex-col items-center space-y-2">
                    <Button
                      size="lg"
                      variant="ghost"
                      className={`w-12 h-12 rounded-full p-0 transition-smooth hover:scale-110 ${
                        isPostBookmarked(post.slug)
                          ? 'text-[#25F4EE] glow-cyan'
                          : 'text-white hover:text-[#25F4EE] hover:bg-white/10'
                      }`}
                      onClick={() => handleBookmark(post.slug)}
                    >
                      <Bookmark className={`h-7 w-7 ${isPostBookmarked(post.slug) ? 'fill-current' : ''}`} />
                    </Button>
                  </div>

                  {/* Follow button */}
                  {post.author_id && (
                    <div className="flex flex-col items-center space-y-2">
                      <FollowButton
                        authorId={post.author_id}
                        authorName={post.author}
                        onLoginRequired={() => setShowLoginModal(true)}
                        variant="icon"
                      />
                    </div>
                  )}

                  {/* Share button */}
                  <div className="flex flex-col items-center space-y-2">
                    <Button
                      size="lg"
                      variant="ghost"
                      className="w-12 h-12 rounded-full p-0 text-white hover:text-[#25F4EE] hover:bg-white/10 transition-smooth hover:scale-110"
                    >
                      <Share className="h-7 w-7" />
                    </Button>
                  </div>
                </div>
              </div>
              );
            })}

            {/* Post progress indicator */}
            <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex items-center space-x-1 z-30">
              <div className="text-[#AFAFAF] caption-small">
                {currentPostIndex + 1} / {posts.length}
                {hasMore && ' + more'}
              </div>
            </div>

            {/* Loading indicator */}
            {loadingMore && (
              <div className="absolute top-4 right-4 z-30">
                <div className="text-white text-sm bg-black/50 px-3 py-1 rounded-full">
                  Loading...
                </div>
              </div>
            )}
          </>
        )}
      </div>

      {/* Audio Progress Bar - only show if current post has audio */}
      {posts[currentPostIndex]?.summary_audio_url && (
        <div className="fixed bottom-0 left-0 right-0 z-40 bg-black/80 backdrop-blur-sm border-t border-white/10">
          {/* Play/Pause Button - centered above progress bar */}
          <div className="flex justify-center pt-4 pb-2">
            <div className="relative">
              <Button
                size="lg"
                variant="ghost"
                className="w-12 h-12 rounded-full audio-play-button bg-black/50 border border-white/10 text-white"
                onClick={toggleAudio}
              >
                {isPlaying ? (
                  <Pause className="h-5 w-5" />
                ) : (
                  <Play className="h-5 w-5 ml-0.5" />
                )}
              </Button>
              
              {/* User interaction required prompt */}
              {showPlayPrompt && (
                <div className="absolute left-1/2 -translate-x-1/2 bottom-16 bg-black/90 backdrop-blur-sm text-white text-xs px-3 py-2 rounded-full whitespace-nowrap border border-white/10">
                  <div className="flex items-center">
                    Tap play button to enable audio
                  </div>
                </div>
              )}
            </div>
          </div>
          
          <div className="px-4 pb-3">
            {/* Progress Bar */}
            <div className="relative">
              <div 
                className="h-1 audio-progress-bar rounded-full cursor-pointer"
                onClick={(e) => {
                  const rect = e.currentTarget.getBoundingClientRect();
                  const percent = (e.clientX - rect.left) / rect.width;
                  const newTime = percent * duration;
                  seekTo(newTime);
                }}
              >
                <div 
                  className="h-1 audio-progress-fill rounded-full transition-all duration-100"
                  style={{ width: `${duration > 0 ? (currentTime / duration) * 100 : 0}%` }}
                />
              </div>
              
              {/* Time Display */}
              <div className="flex justify-between items-center mt-2 text-xs text-[#AFAFAF]">
                <span>{Math.floor(currentTime / 60)}:{Math.floor(currentTime % 60).toString().padStart(2, '0')}</span>
                <span>{Math.floor(duration / 60)}:{Math.floor(duration % 60).toString().padStart(2, '0')}</span>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Login Modal */}
      <LoginModal isOpen={showLoginModal} onClose={() => setShowLoginModal(false)} />
    </div>
  );
}
