'use client';

import { Send } from 'lucide-react';
import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { useAuthStore } from '@/store/authStore';
import { useToast } from '@/components/ui/toast';
import { ApiError, type CommentInfo } from '@/lib/api-client';

interface CommentFormProps {
  postSlug: string;
  parentId?: string;
  onCommentAdded?: (comment: CommentInfo) => void;
  onCancel?: () => void;
  placeholder?: string;
  compact?: boolean;
}

export default function CommentForm({
  postSlug,
  parentId,
  onCommentAdded,
  onCancel,
  placeholder = 'Add a comment...',
  compact = false,
}: CommentFormProps) {
  const [comment, setComment] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { isAuthenticated, getApiClient } = useAuthStore();
  const { showToast } = useToast();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!comment.trim()) {
      showToast('Please enter a comment');
      return;
    }

    if (!isAuthenticated) {
      showToast('Please sign in to comment');
      return;
    }

    try {
      setIsSubmitting(true);
      const apiClient = getApiClient();
      const response = await apiClient.createComment(postSlug, {
        body: comment.trim(),
        parent_id: parentId,
      });

      // Convert response to CommentInfo format
      const newComment: CommentInfo = {
        id: response.id,
        body: response.body,
        author: response.author,
        created_at: response.created_at,
        parent_id: response.parent_id,
      };

      onCommentAdded?.(newComment);
      setComment('');
      showToast('Comment added successfully!', 'success');
      
      if (onCancel) {
        onCancel();
      }
    } catch (err) {
      console.error('Failed to create comment:', err);
      if (err instanceof ApiError) {
        showToast(err.message);
      } else {
        showToast('Failed to add comment');
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  if (!isAuthenticated) {
    return null;
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-3">
      <div className={`relative ${compact ? 'mb-2' : 'mb-4'}`}>
        <textarea
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          placeholder={placeholder}
          disabled={isSubmitting}
          className={`w-full bg-[#121212] border border-white/10 rounded-lg text-white placeholder:text-[#AFAFAF] resize-none focus:outline-none focus:border-[#FE2C55]/50 focus:ring-1 focus:ring-[#FE2C55]/20 transition-colors ${
            compact ? 'p-3 text-sm min-h-[80px]' : 'p-4 text-base min-h-[100px]'
          }`}
          rows={compact ? 3 : 4}
        />
      </div>

      <div className="flex justify-end space-x-2">
        {onCancel && (
          <Button
            type="button"
            variant="ghost"
            size="sm"
            onClick={onCancel}
            disabled={isSubmitting}
            className="text-[#AFAFAF] hover:text-white hover:bg-white/10"
          >
            Cancel
          </Button>
        )}
        <Button
          type="submit"
          size="sm"
          disabled={!comment.trim() || isSubmitting}
          className="bg-[#FE2C55] text-white hover:bg-[#FE2C55]/90 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {isSubmitting ? (
            'Posting...'
          ) : (
            <>
              <Send className="h-4 w-4 mr-2" />
              {parentId ? 'Reply' : 'Comment'}
            </>
          )}
        </Button>
      </div>
    </form>
  );
}