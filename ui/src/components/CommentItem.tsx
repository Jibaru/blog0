'use client';

import { MessageCircle, MoreHorizontal } from 'lucide-react';
import { useState } from 'react';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { type CommentInfo } from '@/lib/api-client';
import CommentForm from '@/components/CommentForm';

interface CommentItemProps {
  comment: CommentInfo;
  postSlug: string;
  onReplyAdded?: (reply: CommentInfo) => void;
  isReply?: boolean;
}

export default function CommentItem({ 
  comment, 
  postSlug, 
  onReplyAdded, 
  isReply = false 
}: CommentItemProps) {
  const [showReplyForm, setShowReplyForm] = useState(false);

  const handleReplyAdded = (reply: CommentInfo) => {
    onReplyAdded?.(reply);
    setShowReplyForm(false);
  };

  return (
    <div className={`${isReply ? 'ml-6 pl-4 border-l border-white/10' : ''}`}>
      <div className="flex items-start space-x-3">
        <Avatar className={`border border-white/10 mt-1 ${isReply ? 'w-8 h-8' : 'w-10 h-10'}`}>
          <AvatarFallback className="bg-[#121212] text-[#AFAFAF] text-sm">
            {comment.author.name.charAt(0).toUpperCase()}
          </AvatarFallback>
        </Avatar>
        
        <div className="flex-1 space-y-2">
          <div>
            <span className={`text-white font-medium ${isReply ? 'text-sm' : 'text-sm'} mr-2`}>
              {comment.author.name}
            </span>
            <span className="text-[#AFAFAF] caption-small">
              {new Date(comment.created_at).toLocaleString()}
            </span>
          </div>
          
          <p className={`text-white leading-relaxed ${isReply ? 'text-sm' : 'body-medium'}`}>
            {comment.body}
          </p>
          
          <div className="flex items-center space-x-4">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => setShowReplyForm(!showReplyForm)}
              className="text-[#AFAFAF] hover:text-white hover:bg-white/5 p-0 h-auto font-medium text-xs"
            >
              <MessageCircle className="h-3 w-3 mr-1" />
              Reply
            </Button>
            
            <Button
              variant="ghost"
              size="sm"
              className="text-[#AFAFAF] hover:text-white hover:bg-white/5 p-1 h-auto"
            >
              <MoreHorizontal className="h-3 w-3" />
            </Button>
          </div>
          
          {showReplyForm && (
            <div className="pt-3">
              <CommentForm
                postSlug={postSlug}
                parentId={comment.id}
                onCommentAdded={handleReplyAdded}
                onCancel={() => setShowReplyForm(false)}
                placeholder={`Reply to ${comment.author.name}...`}
                compact
              />
            </div>
          )}
        </div>
      </div>
    </div>
  );
}