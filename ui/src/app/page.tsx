'use client';

import { Bookmark, Calendar, Heart, LogIn, MessageCircle, Share, User, Plus } from 'lucide-react';
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
  const [currentPage, setCurrentPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const loadingRef = useRef(false);
  const observerRef = useRef<IntersectionObserver | null>(null);
  const lastPostRef = useRef<HTMLDivElement>(null);
  
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

  const fetchPosts = async (page: number, append = false) => {
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
  };

  const loadMore = useCallback(() => {
    if (!loadingRef.current && hasMore && !loadingMore) {
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
  }, [hasMore, loadingMore, currentPage]);

  const handleLike = async (slug: string) => {
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
  }, []);

  // Set up intersection observer when posts change
  useEffect(() => {
    if (!lastPostRef.current || !hasMore) return;

    // Clean up previous observer
    if (observerRef.current) {
      observerRef.current.disconnect();
    }

    // Create new observer
    observerRef.current = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting && hasMore && !loadingMore && !loadingRef.current) {
          console.log('Last post is visible, loading more posts...');
          loadMore();
        }
      },
      {
        rootMargin: '200px',
        threshold: 0.1,
      }
    );

    // Observe the last post
    observerRef.current.observe(lastPostRef.current);

    // Cleanup
    return () => {
      if (observerRef.current) {
        observerRef.current.disconnect();
      }
    };
  }, [posts, hasMore, loadingMore, loadMore]);

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

      {/* Posts container */}
      <div className="posts-scroll">
        {posts.length === 0 ? (
          <div className="post-item">
            <div className="text-center space-y-4">
              <div className="text-2xl text-[#AFAFAF]">No posts available</div>
              <div className="accent-text text-lg">Check back later!</div>
            </div>
          </div>
        ) : (
          <>
            <div className="space-y-0">
              {posts.map((post, index) => (
                <div 
                  key={post.slug} 
                  ref={index === posts.length - 1 ? lastPostRef : null}
                  className="post-item relative"
                >
                {/* Background content area */}
                <div className="absolute inset-0 bg-gradient-to-b from-black/0 via-black/20 to-black/80" />

                {/* Main content */}
                <div className="relative z-10 w-full max-w-sm mx-auto px-4 h-full flex flex-col justify-between">
                  {/* Author info at top */}
                  <div className="pt-20 flex items-center justify-between pr-16">
                    <div className="flex items-center space-x-3">
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
                  </div>

                  {/* Bottom overlay with post info */}
                  <div className="pb-8 pr-16">
                    <Link href={`/post/${post.slug}`} className="block">
                      <div className="space-y-4">
                        {/* Post title */}
                        <h2 className="text-xl md:text-2xl font-bold leading-tight text-white break-words">
                          {post.title}
                        </h2>

                        {/* Post metadata */}
                        <div className="flex items-center space-x-4 text-[#AFAFAF]">
                          {post.published_at && (
                            <div className="flex items-center caption-small">
                              <Calendar className="h-3 w-3 mr-1" />
                              {new Date(post.published_at).toLocaleDateString()}
                            </div>
                          )}
                          <div className="caption-small text-[#FE2C55]">#blog #tech</div>
                        </div>

                        {/* Engagement preview */}
                        <div className="flex items-center space-x-4 text-[#AFAFAF] caption-small">
                          <span className="flex items-center">
                            <Heart className="h-3 w-3 mr-1" />
                            {postLikeCounts[post.slug] ?? post.like_count} likes
                          </span>
                          <span className="flex items-center">
                            <MessageCircle className="h-3 w-3 mr-1" />
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
              ))}
            </div>

            {/* Loading more indicator */}
            {loadingMore && (
              <div className="flex items-center justify-center py-8">
                <div className="text-white text-lg">Loading more posts...</div>
              </div>
            )}

            {/* End of posts indicator */}
            {!hasMore && posts.length > 0 && (
              <div className="flex items-center justify-center py-8">
                <div className="text-[#AFAFAF] text-sm">You&apos;ve reached the end!</div>
              </div>
            )}
          </>
        )}
      </div>

      {/* Login Modal */}
      <LoginModal isOpen={showLoginModal} onClose={() => setShowLoginModal(false)} />
    </div>
  );
}
