'use client';

import { Bookmark, Calendar, Heart, LogIn, MessageCircle, Share, User } from 'lucide-react';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import LoginModal from '@/components/auth/LoginModal';
import UserMenu from '@/components/auth/UserMenu';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import Blog0ApiClient, { type PostItem } from '@/lib/api-client';
import { useAuthStore } from '@/store/authStore';

const api = new Blog0ApiClient();

export default function Home() {
  const [posts, setPosts] = useState<PostItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showLoginModal, setShowLoginModal] = useState(false);

  const { isAuthenticated } = useAuthStore();

  useEffect(() => {
    async function fetchPosts() {
      try {
        const response = await api.listPosts({ page: 1, per_page: 20 });
        setPosts(response.items);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to fetch posts');
      } finally {
        setLoading(false);
      }
    }

    fetchPosts();
  }, []);

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
          <div className="space-y-0">
            {posts.map((post) => (
              <div key={post.slug} className="post-item relative">
                {/* Background content area */}
                <div className="absolute inset-0 bg-gradient-to-b from-black/0 via-black/20 to-black/80" />

                {/* Main content */}
                <div className="relative z-10 w-full max-w-sm mx-auto px-4 h-full flex flex-col justify-between">
                  {/* Author info at top */}
                  <div className="pt-20 flex items-center space-x-3">
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
                            {Math.floor(Math.random() * 1000)} likes
                          </span>
                          <span className="flex items-center">
                            <MessageCircle className="h-3 w-3 mr-1" />
                            {Math.floor(Math.random() * 50)} comments
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
                      className="w-12 h-12 rounded-full p-0 text-white hover:text-[#FE2C55] hover:bg-white/10 transition-smooth hover:scale-110"
                      onClick={() => (isAuthenticated ? null : setShowLoginModal(true))}
                    >
                      <Heart className="h-7 w-7" />
                    </Button>
                    <span className="caption-small text-white font-medium">
                      {Math.floor(Math.random() * 1000)}
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
                      {Math.floor(Math.random() * 50)}
                    </span>
                  </div>

                  {/* Bookmark button */}
                  <div className="flex flex-col items-center space-y-2">
                    <Button
                      size="lg"
                      variant="ghost"
                      className="w-12 h-12 rounded-full p-0 text-white hover:text-[#FE2C55] hover:bg-white/10 transition-smooth hover:scale-110"
                      onClick={() => (isAuthenticated ? null : setShowLoginModal(true))}
                    >
                      <Bookmark className="h-7 w-7" />
                    </Button>
                  </div>

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
        )}
      </div>

      {/* Login Modal */}
      <LoginModal isOpen={showLoginModal} onClose={() => setShowLoginModal(false)} />
    </div>
  );
}
