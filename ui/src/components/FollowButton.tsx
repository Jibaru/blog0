'use client';

import { UserPlus, UserMinus } from 'lucide-react';
import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { useAuthStore } from '@/store/authStore';
import { useToast } from '@/components/ui/toast';
import { ApiError } from '@/lib/api-client';

interface FollowButtonProps {
  authorId: string;
  authorName: string;
  initialFollowing?: boolean;
  initialFollowersCount?: number;
  onFollowChange?: (following: boolean, followersCount: number) => void;
  onLoginRequired?: () => void;
  variant?: 'default' | 'compact' | 'icon';
  className?: string;
}

export default function FollowButton({
  authorId,
  authorName,
  initialFollowing = false,
  initialFollowersCount = 0,
  onFollowChange,
  onLoginRequired,
  variant = 'default',
  className = '',
}: FollowButtonProps) {
  const [following, setFollowing] = useState(initialFollowing);
  const [, setFollowersCount] = useState(initialFollowersCount);
  const [isLoading, setIsLoading] = useState(false);
  const { isAuthenticated, getApiClient, user, isUserFollowed, updateUserFollow } = useAuthStore();
  
  // Use profile data if available, otherwise fall back to initial state
  const isFollowingAuthor = isUserFollowed(authorId) || following;
  const { showToast } = useToast();

  // Don't show follow button if no author ID is provided
  if (!authorId) {
    return null;
  }

  // Don't show follow button for own profile
  if (user && user.id === authorId) {
    return null;
  }

  const handleFollow = async () => {
    if (!isAuthenticated) {
      onLoginRequired?.();
      return;
    }

    try {
      setIsLoading(true);
      const apiClient = getApiClient();
      
      let response;
      if (isFollowingAuthor) {
        response = await apiClient.unfollowUser(authorId);
        showToast(`Unfollowed ${authorName}`, 'success');
      } else {
        response = await apiClient.followUser(authorId);
        showToast(`Following ${authorName}`, 'success');
      }
      
      // Update local state
      setFollowing(response.following);
      setFollowersCount(response.followers_count);
      
      // Update profile state
      updateUserFollow(authorId, response.following);
      
      onFollowChange?.(response.following, response.followers_count);
      
    } catch (err) {
      console.error('Failed to toggle follow:', err);
      if (err instanceof ApiError) {
        showToast(err.message);
      } else {
        showToast(following ? 'Failed to unfollow' : 'Failed to follow');
      }
    } finally {
      setIsLoading(false);
    }
  };

  if (variant === 'icon') {
    return (
      <Button
        size="sm"
        variant="ghost"
        onClick={handleFollow}
        disabled={isLoading}
        className={`w-10 h-10 rounded-full p-0 transition-smooth hover:scale-110 ${
          isFollowingAuthor
            ? 'text-[#25F4EE] hover:text-[#25F4EE]/80'
            : 'text-white hover:text-[#25F4EE] hover:bg-white/10'
        } ${className}`}
      >
        {isFollowingAuthor ? (
          <UserMinus className="h-5 w-5" />
        ) : (
          <UserPlus className="h-5 w-5" />
        )}
      </Button>
    );
  }

  if (variant === 'compact') {
    return (
      <Button
        size="sm"
        onClick={handleFollow}
        disabled={isLoading}
        className={`${
          isFollowingAuthor
            ? 'bg-transparent border border-[#25F4EE] text-[#25F4EE] hover:bg-[#25F4EE]/10'
            : 'bg-[#25F4EE] text-black hover:bg-[#25F4EE]/90'
        } transition-smooth ${className}`}
      >
        {isLoading ? (
          'Loading...'
        ) : isFollowingAuthor ? (
          <>
            <UserMinus className="h-3 w-3 mr-1" />
            Following
          </>
        ) : (
          <>
            <UserPlus className="h-3 w-3 mr-1" />
            Follow
          </>
        )}
      </Button>
    );
  }

  return (
    <Button
      onClick={handleFollow}
      disabled={isLoading}
      className={`${
        isFollowingAuthor
          ? 'bg-transparent border border-[#25F4EE] text-[#25F4EE] hover:bg-[#25F4EE]/10'
          : 'bg-[#25F4EE] text-black hover:bg-[#25F4EE]/90'
      } transition-smooth ${className}`}
    >
      {isLoading ? (
        'Loading...'
      ) : isFollowingAuthor ? (
        <>
          <UserMinus className="h-4 w-4 mr-2" />
          Following
        </>
      ) : (
        <>
          <UserPlus className="h-4 w-4 mr-2" />
          Follow
        </>
      )}
    </Button>
  );
}