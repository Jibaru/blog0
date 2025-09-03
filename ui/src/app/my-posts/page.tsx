'use client';

import { ArrowLeft, Calendar, Edit3, Eye, Trash2, Plus } from 'lucide-react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { useEffect, useState, useCallback } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { useAuthStore } from '@/store/authStore';
import { useToast } from '@/components/ui/toast';
import { ApiError, type MyPostItem } from '@/lib/api-client';

export default function MyPostsPage() {
  const [posts, setPosts] = useState<MyPostItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [deletingPostId, setDeletingPostId] = useState<string | null>(null);
  
  const router = useRouter();
  const { isAuthenticated, getApiClient } = useAuthStore();
  const { showToast } = useToast();

  // Redirect if not authenticated
  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/');
      return;
    }
  }, [isAuthenticated, router]);

  const fetchMyPosts = useCallback(async () => {
    try {
      setLoading(true);
      const apiClient = getApiClient();
      const response = await apiClient.listMyPosts({ per_page: 50, order: 'DESC' });
      setPosts(response.items);
    } catch (err) {
      const errorMessage = err instanceof ApiError ? err.message : 'Failed to fetch posts';
      setError(errorMessage);
      showToast(errorMessage);
    } finally {
      setLoading(false);
    }
  }, [getApiClient, showToast]);

  useEffect(() => {
    if (isAuthenticated) {
      fetchMyPosts();
    }
  }, [isAuthenticated, fetchMyPosts]);

  const handleDeletePost = async (postId: string, postTitle: string) => {
    if (!confirm(`Are you sure you want to delete "${postTitle}"? This action cannot be undone.`)) {
      return;
    }

    try {
      setDeletingPostId(postId);
      const apiClient = getApiClient();
      const post = posts.find(p => p.id === postId);
      if (!post) return;

      await apiClient.deletePost(post.slug);
      setPosts(prev => prev.filter(p => p.id !== postId));
      showToast('Post deleted successfully', 'success');
    } catch (err) {
      console.error('Failed to delete post:', err);
      if (err instanceof ApiError) {
        showToast(err.message);
      } else {
        showToast('Failed to delete post');
      }
    } finally {
      setDeletingPostId(null);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status.toLowerCase()) {
      case 'published':
        return 'bg-[#25F4EE] text-black';
      case 'draft':
        return 'bg-[#AFAFAF] text-black';
      default:
        return 'bg-[#AFAFAF] text-black';
    }
  };

  if (!isAuthenticated) {
    return null;
  }

  if (loading) {
    return (
      <div className="min-h-screen bg-black mesh-background flex items-center justify-center">
        <div className="text-2xl accent-text font-bold">Loading your posts...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-black mesh-background flex items-center justify-center">
        <div className="text-center space-y-4">
          <div className="text-xl text-[#FE2C55]">Error: {error}</div>
          <Button
            onClick={fetchMyPosts}
            className="bg-[#FE2C55] text-white border-0 hover:bg-[#FE2C55]/80"
          >
            Retry
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-black mesh-background">
      {/* Header */}
      <header className="sticky top-0 z-50 bg-black/80 backdrop-blur-sm border-b border-white/10">
        <div className="max-w-6xl mx-auto px-4 h-16 flex items-center justify-between">
          <div className="flex items-center space-x-4">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => router.back()}
              className="text-white hover:text-[#FE2C55] hover:bg-white/10"
            >
              <ArrowLeft className="h-4 w-4 mr-2" />
              Back
            </Button>
            <h1 className="text-white font-bold text-xl">My Posts</h1>
          </div>

          <Link href="/create">
            <Button
              size="sm"
              className="bg-[#FE2C55] text-white hover:bg-[#FE2C55]/90"
            >
              <Plus className="h-4 w-4 mr-2" />
              New Post
            </Button>
          </Link>
        </div>
      </header>

      {/* Content */}
      <div className="max-w-6xl mx-auto px-4 py-8">
        {posts.length === 0 ? (
          <div className="text-center py-16">
            <div className="text-2xl text-[#AFAFAF] mb-4">No posts yet</div>
            <div className="text-[#AFAFAF] mb-6">Start creating your first blog post!</div>
            <Link href="/create">
              <Button className="bg-[#FE2C55] text-white hover:bg-[#FE2C55]/90">
                <Plus className="h-4 w-4 mr-2" />
                Create Your First Post
              </Button>
            </Link>
          </div>
        ) : (
          <>
            {/* Stats */}
            <div className="mb-8">
              <div className="text-[#AFAFAF] text-sm">
                {posts.length} post{posts.length !== 1 ? 's' : ''} •{' '}
                {posts.filter(p => p.status.toLowerCase() === 'published').length} published •{' '}
                {posts.filter(p => p.status.toLowerCase() === 'draft').length} draft{posts.filter(p => p.status.toLowerCase() === 'draft').length !== 1 ? 's' : ''}
              </div>
            </div>

            {/* Posts Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {posts.map((post) => (
                <Card key={post.id} className="bg-[#121212] border-white/10 hover:border-white/20 transition-colors">
                  <CardContent className="p-6">
                    {/* Status Badge */}
                    <div className="flex items-center justify-between mb-4">
                      <Badge className={`${getStatusColor(post.status)} font-medium text-xs`}>
                        {post.status}
                      </Badge>
                      <div className="text-[#AFAFAF] text-xs flex items-center">
                        <Calendar className="h-3 w-3 mr-1" />
                        {new Date(post.updated_at).toLocaleDateString()}
                      </div>
                    </div>

                    {/* Post Title */}
                    <h3 className="text-white font-bold text-lg mb-3 line-clamp-2 leading-tight">
                      {post.title}
                    </h3>

                    {/* Post Meta */}
                    <div className="text-[#AFAFAF] text-sm mb-4 space-y-1">
                      <div>Created: {new Date(post.created_at).toLocaleDateString()}</div>
                      {post.published_at && (
                        <div>Published: {new Date(post.published_at).toLocaleDateString()}</div>
                      )}
                    </div>

                    {/* Actions */}
                    <div className="flex items-center space-x-2">
                      {post.status.toLowerCase() === 'published' && (
                        <Link href={`/post/${post.slug}`} className="flex-1">
                          <Button
                            variant="outline"
                            size="sm"
                            className="w-full border-white/10 text-[#AFAFAF] hover:text-white hover:bg-white/5"
                          >
                            <Eye className="h-3 w-3 mr-2" />
                            View
                          </Button>
                        </Link>
                      )}
                      
                      <Link href={`/edit/${post.slug}`} className="flex-1">
                        <Button
                          size="sm"
                          className="w-full bg-[#25F4EE] text-black hover:bg-[#25F4EE]/90"
                        >
                          <Edit3 className="h-3 w-3 mr-2" />
                          Edit
                        </Button>
                      </Link>

                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => handleDeletePost(post.id, post.title)}
                        disabled={deletingPostId === post.id}
                        className="border-[#FE2C55]/20 text-[#FE2C55] hover:bg-[#FE2C55]/10 hover:border-[#FE2C55]/40"
                      >
                        {deletingPostId === post.id ? (
                          <div className="w-3 h-3" />
                        ) : (
                          <Trash2 className="h-3 w-3" />
                        )}
                      </Button>
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          </>
        )}
      </div>
    </div>
  );
}