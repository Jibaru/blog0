'use client';

import {
  ArrowLeft,
  Bookmark,
  Calendar,
  Heart,
  LogIn,
  MessageCircle,
  Share,
  User,
} from 'lucide-react';
import { useParams, useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import LoginModal from '@/components/auth/LoginModal';
import UserMenu from '@/components/auth/UserMenu';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import Blog0ApiClient, { type GetPostBySlugResp } from '@/lib/api-client';
import { useAuthStore } from '@/store/authStore';

const api = new Blog0ApiClient();

export default function PostPage() {
  const params = useParams();
  const router = useRouter();
  const slug = params.slug as string;

  const [post, setPost] = useState<GetPostBySlugResp | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [liked, setLiked] = useState(false);
  const [likesCount, setLikesCount] = useState(0);
  const [bookmarked, setBookmarked] = useState(false);
  const [showLoginModal, setShowLoginModal] = useState(false);

  const { isAuthenticated } = useAuthStore();

  useEffect(() => {
    if (!slug) return;

    async function fetchPost() {
      try {
        const response = await api.getPostBySlug(slug);
        setPost(response);
        setLikesCount(response.likes_count);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to fetch post');
      } finally {
        setLoading(false);
      }
    }

    fetchPost();
  }, [slug]);

  const handleLike = async () => {
    if (!post) return;

    if (!isAuthenticated) {
      setShowLoginModal(true);
      return;
    }

    try {
      const response = await api.toggleLike(post.slug);
      setLiked(response.liked);
      setLikesCount(response.likes_count);
    } catch (err) {
      console.error('Failed to toggle like:', err);
    }
  };

  const handleBookmark = async () => {
    if (!post) return;

    if (!isAuthenticated) {
      setShowLoginModal(true);
      return;
    }

    try {
      if (bookmarked) {
        await api.unbookmarkPost(post.slug);
        setBookmarked(false);
      } else {
        await api.bookmarkPost(post.slug);
        setBookmarked(true);
      }
    } catch (err) {
      console.error('Failed to toggle bookmark:', err);
    }
  };

  const handleShare = async () => {
    if (navigator.share) {
      try {
        await navigator.share({
          title: post?.title,
          url: window.location.href,
        });
      } catch (_err) {
        console.log('Share cancelled');
      }
    } else {
      navigator.clipboard.writeText(window.location.href);
      alert('Link copied to clipboard!');
    }
  };

  if (loading) {
    return (
      <div className="post-container flex items-center justify-center">
        <div className="text-2xl accent-text font-bold">Loading post...</div>
      </div>
    );
  }

  if (error || !post) {
    return (
      <div className="post-container flex flex-col items-center justify-center space-y-6">
        <div className="text-2xl text-[#FE2C55] text-center">{error || 'Post not found'}</div>
        <Button
          onClick={() => router.back()}
          className="bg-[#FE2C55] text-white border-0 hover:bg-[#FE2C55]/80 px-8 py-3"
        >
          <ArrowLeft className="h-4 w-4 mr-2" />
          Go Back
        </Button>
      </div>
    );
  }

  return (
    <div className="bg-black min-h-screen mesh-background">
      {/* Slim top bar with back button */}
      <header className="fixed top-0 left-0 right-0 z-50 h-16 flex items-center justify-between px-4 bg-black/80 backdrop-blur-sm">
        <div className="flex items-center">
          <Button
            variant="ghost"
            size="sm"
            onClick={() => router.back()}
            className="text-white hover:text-[#FE2C55] hover:bg-white/10 mr-4 p-2"
          >
            <ArrowLeft className="h-5 w-5" />
          </Button>
          <h1 className="accent-text text-lg font-bold">Blog0</h1>
        </div>

        {/* Auth section */}
        <div className="flex items-center">
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

      {/* Main content container */}
      <div className="pt-16 relative">
        <div className="max-w-md mx-auto relative min-h-screen">
          {/* Content area with generous padding */}
          <div className="px-6 pb-32">
            {/* Author section */}
            <div className="pt-6 pb-6 flex items-center space-x-4">
              <Avatar className="w-16 h-16 border-2 border-white/20">
                <AvatarFallback className="bg-[#FE2C55] text-white font-bold text-xl">
                  {post.author.name.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <div className="space-y-1">
                <div className="text-white font-bold text-lg">{post.author.name}</div>
                <div className="text-[#AFAFAF] caption-small flex items-center">
                  <User className="h-3 w-3 mr-1" />@{post.author.name}
                </div>
                {post.published_at && (
                  <div className="text-[#AFAFAF] caption-small flex items-center">
                    <Calendar className="h-3 w-3 mr-1" />
                    {new Date(post.published_at).toLocaleString()}
                  </div>
                )}
              </div>
            </div>

            {/* Post title */}
            <div className="pb-8">
              <h1 className="title-large text-white leading-tight mb-4">{post.title}</h1>

              {/* Hashtags */}
              <div className="flex flex-wrap gap-2">
                <span className="text-[#AFAFAF] caption-small font-medium">#blog</span>
                <span className="text-[#AFAFAF] caption-small font-medium">#tech</span>
                <span className="text-[#AFAFAF] caption-small font-medium">#development</span>
              </div>
            </div>

            {/* Post content */}
            <div className="pb-8">
              <div className="body-medium text-white whitespace-pre-wrap leading-relaxed">
                {post.raw_markdown}
              </div>
            </div>

            {/* Engagement stats */}
            <div className="pb-6">
              <div className="flex items-center space-x-6 text-[#AFAFAF] caption-small">
                <span className="flex items-center">
                  <Heart className="h-4 w-4 mr-2" />
                  {likesCount} likes
                </span>
                <span className="flex items-center">
                  <MessageCircle className="h-4 w-4 mr-2" />
                  {post.comments.length} comments
                </span>
                <span className="flex items-center">
                  <Share className="h-4 w-4 mr-2" />
                  Share
                </span>
              </div>
            </div>

            {/* Comments section */}
            {post.comments.length > 0 && (
              <div className="space-y-6">
                <Separator className="bg-[#121212]" />

                <div className="space-y-1">
                  <h3 className="text-white font-bold text-lg mb-4">Comments</h3>
                  <div className="text-[#AFAFAF] caption-small mb-6">
                    {post.comments.length} comment{post.comments.length !== 1 ? 's' : ''}
                  </div>
                </div>

                <div className="space-y-6">
                  {post.comments.map((comment) => (
                    <div key={comment.id} className="flex items-start space-x-3">
                      <Avatar className="w-10 h-10 border border-white/10 mt-1">
                        <AvatarFallback className="bg-[#121212] text-[#AFAFAF] text-sm">
                          {comment.author.name.charAt(0).toUpperCase()}
                        </AvatarFallback>
                      </Avatar>
                      <div className="flex-1 space-y-2">
                        <div>
                          <span className="text-white font-medium text-sm mr-2">
                            {comment.author.name}
                          </span>
                          <span className="text-[#AFAFAF] caption-small">
                            {new Date(comment.created_at).toLocaleString()}
                          </span>
                        </div>
                        <p className="body-medium text-white leading-relaxed">{comment.body}</p>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>

          {/* Fixed right sidebar with action buttons */}
          <div className="fixed right-4 bottom-32 flex flex-col space-y-8 z-20">
            {/* Like button */}
            <div className="flex flex-col items-center space-y-2">
              <Button
                size="lg"
                variant="ghost"
                className={`w-14 h-14 rounded-full p-0 transition-smooth hover:scale-110 ${
                  liked
                    ? 'text-[#FE2C55] glow-accent'
                    : 'text-white hover:text-[#FE2C55] hover:bg-white/10'
                }`}
                onClick={handleLike}
              >
                <Heart className={`h-8 w-8 ${liked ? 'fill-current' : ''}`} />
              </Button>
              <span className="caption-small text-white font-bold">{likesCount}</span>
            </div>

            {/* Comment button */}
            <div className="flex flex-col items-center space-y-2">
              <Button
                size="lg"
                variant="ghost"
                className="w-14 h-14 rounded-full p-0 text-white hover:text-[#25F4EE] hover:bg-white/10 transition-smooth hover:scale-110"
              >
                <MessageCircle className="h-8 w-8" />
              </Button>
              <span className="caption-small text-white font-bold">{post.comments.length}</span>
            </div>

            {/* Bookmark button */}
            <div className="flex flex-col items-center space-y-2">
              <Button
                size="lg"
                variant="ghost"
                className={`w-14 h-14 rounded-full p-0 transition-smooth hover:scale-110 ${
                  bookmarked
                    ? 'text-[#25F4EE] glow-cyan'
                    : 'text-white hover:text-[#25F4EE] hover:bg-white/10'
                }`}
                onClick={handleBookmark}
              >
                <Bookmark className={`h-8 w-8 ${bookmarked ? 'fill-current' : ''}`} />
              </Button>
            </div>

            {/* Share button */}
            <div className="flex flex-col items-center space-y-2">
              <Button
                size="lg"
                variant="ghost"
                className="w-14 h-14 rounded-full p-0 text-white hover:text-[#FE2C55] hover:bg-white/10 transition-smooth hover:scale-110"
                onClick={handleShare}
              >
                <Share className="h-8 w-8" />
              </Button>
            </div>
          </div>
        </div>
      </div>

      {/* Login Modal */}
      <LoginModal isOpen={showLoginModal} onClose={() => setShowLoginModal(false)} />
    </div>
  );
}
