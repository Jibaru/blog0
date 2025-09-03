'use client';

import MDEditor from '@uiw/react-md-editor';
import { ArrowLeft, Save, Eye, Edit3 } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { useAuthStore } from '@/store/authStore';
import { useToast } from '@/components/ui/toast';
import { ApiError } from '@/lib/api-client';

interface PostEditorProps {
  postSlug?: string; // If provided, we're editing an existing post
}

export default function PostEditor({ postSlug }: PostEditorProps) {
  const [title, setTitle] = useState('');
  const [slug, setSlug] = useState('');
  const [content, setContent] = useState('');
  const [isPublished, setIsPublished] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [isLoadingPost, setIsLoadingPost] = useState(!!postSlug);
  const [previewMode, setPreviewMode] = useState(false);

  const router = useRouter();
  const { isAuthenticated, getApiClient } = useAuthStore();
  const { showToast } = useToast();

  const isEditing = !!postSlug;

  // Auto-generate slug from title
  const generateSlug = (title: string) => {
    return title
      .toLowerCase()
      .trim()
      .replace(/[^a-z0-9\s-]/g, '')
      .replace(/\s+/g, '-')
      .replace(/-+/g, '-')
      .replace(/^-|-$/g, '');
  };

  // Update slug when title changes (only for new posts)
  useEffect(() => {
    if (!isEditing && title) {
      setSlug(generateSlug(title));
    }
  }, [title, isEditing]);

  // Load existing post if editing
  useEffect(() => {
    if (!postSlug || !isAuthenticated) return;

    const loadPost = async () => {
      try {
        setIsLoadingPost(true);
        const apiClient = getApiClient();
        const post = await apiClient.getPostBySlug(postSlug);
        
        setTitle(post.title);
        setSlug(post.slug);
        setContent(post.raw_markdown);
        setIsPublished(!!post.published_at);
      } catch (err) {
        console.error('Failed to load post:', err);
        if (err instanceof ApiError) {
          showToast(err.message);
        } else {
          showToast('Failed to load post');
        }
        router.back();
      } finally {
        setIsLoadingPost(false);
      }
    };

    loadPost();
  }, [postSlug, isAuthenticated, getApiClient, showToast, router]);

  // Redirect if not authenticated
  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/');
    }
  }, [isAuthenticated, router]);

  const handleSave = async (publish = false) => {
    if (!title.trim()) {
      showToast('Please enter a post title');
      return;
    }

    if (!content.trim()) {
      showToast('Please enter post content');
      return;
    }

    if (!isEditing && !slug.trim()) {
      showToast('Please enter a post slug');
      return;
    }

    try {
      setIsLoading(true);
      const apiClient = getApiClient();

      if (isEditing) {
        // Update existing post
        await apiClient.updatePost(postSlug, {
          title: title.trim(),
          raw_markdown: content,
          publish: publish || isPublished,
        });
        
        showToast('Post updated successfully!', 'success');
        if (publish && !isPublished) {
          setIsPublished(true);
        }
      } else {
        // Create new post
        await apiClient.createPost({
          title: title.trim(),
          slug: slug.trim(),
          raw_markdown: content,
          publish,
        });
        
        showToast('Post created successfully!', 'success');
        
        // Navigate to the new post
        router.push(`/post/${slug.trim()}`);
        return;
      }
    } catch (err) {
      console.error('Failed to save post:', err);
      if (err instanceof ApiError) {
        showToast(err.message);
      } else {
        showToast('Failed to save post');
      }
    } finally {
      setIsLoading(false);
    }
  };

  if (!isAuthenticated) {
    return null;
  }

  if (isLoadingPost) {
    return (
      <div className="min-h-screen bg-black flex items-center justify-center">
        <div className="text-white text-xl">Loading post...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-black">
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
            <h1 className="text-white font-bold text-lg">
              {isEditing ? 'Edit Post' : 'Create Post'}
            </h1>
          </div>

          <div className="flex items-center space-x-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => setPreviewMode(!previewMode)}
              className="text-white hover:text-[#25F4EE] hover:bg-white/10"
            >
              {previewMode ? <Edit3 className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
              {previewMode ? 'Edit' : 'Preview'}
            </Button>
            
            <Button
              onClick={() => handleSave(false)}
              disabled={isLoading}
              variant="outline"
              size="sm"
              className="border-white/10 text-white hover:bg-white/5"
            >
              <Save className="h-4 w-4 mr-2" />
              {isLoading ? 'Saving...' : 'Save Draft'}
            </Button>

            <Button
              onClick={() => handleSave(true)}
              disabled={isLoading}
              size="sm"
              className="bg-[#FE2C55] text-white hover:bg-[#FE2C55]/90"
            >
              {isLoading ? 'Publishing...' : isPublished ? 'Update' : 'Publish'}
            </Button>
          </div>
        </div>
      </header>

      {/* Content */}
      <div className="max-w-6xl mx-auto p-4">
        <div className="space-y-6">
          {/* Title Input */}
          <div>
            <input
              type="text"
              placeholder="Enter your post title..."
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full bg-transparent border-none text-white text-3xl font-bold placeholder:text-[#AFAFAF] focus:outline-none focus:ring-0 resize-none"
            />
          </div>

          {/* Slug Input (only for new posts) */}
          {!isEditing && (
            <div>
              <label className="block text-[#AFAFAF] text-sm font-medium mb-2">
                URL Slug
              </label>
              <div className="flex items-center">
                <span className="text-[#AFAFAF] mr-2">/post/</span>
                <input
                  type="text"
                  placeholder="post-url-slug"
                  value={slug}
                  onChange={(e) => setSlug(e.target.value)}
                  className="flex-1 bg-[#121212] border border-white/10 rounded px-3 py-2 text-white placeholder:text-[#AFAFAF] focus:outline-none focus:border-[#FE2C55]/50 focus:ring-1 focus:ring-[#FE2C55]/20"
                />
              </div>
            </div>
          )}

          {/* Markdown Editor */}
          <div className="border border-white/10 rounded-lg overflow-hidden">
            <div data-color-mode="dark">
              <MDEditor
                value={content}
                onChange={(val) => setContent(val || '')}
                preview={previewMode ? 'preview' : 'edit'}
                hideToolbar={previewMode}
                height={600}
                data-color-mode="dark"
                style={{
                  backgroundColor: 'transparent',
                }}
                textareaProps={{
                  placeholder: 'Write your post content in Markdown...',
                  style: {
                    fontSize: '14px',
                    lineHeight: '1.6',
                    fontFamily: 'ui-monospace, SFMono-Regular, "SF Mono", Monaco, Inconsolata, "Roboto Mono", monospace',
                    backgroundColor: '#121212',
                    color: '#ffffff',
                    border: 'none',
                    resize: 'none',
                  },
                }}
              />
            </div>
          </div>

          {/* Status Info */}
          {isEditing && (
            <div className="text-center text-[#AFAFAF] text-sm">
              Status: {isPublished ? 'Published' : 'Draft'} • 
              Last updated: {new Date().toLocaleString()}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}