import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import Blog0ApiClient, { type GetProfileResp } from '@/lib/api-client';

export interface User {
  id: string;
  name: string;
  email?: string;
}

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
  // Efficient hash maps for profile data
  likedPosts: Record<string, boolean>; // postId -> true
  bookmarkedPosts: Record<string, boolean>; // postId -> true
  followingUsers: Record<string, boolean>; // userId -> true
}

interface AuthActions {
  login: (provider: string) => Promise<void>;
  logout: () => void;
  setToken: (token: string) => void;
  setUser: (user: User) => void;
  clearError: () => void;
  setLoading: (loading: boolean) => void;
  checkAuth: () => Promise<void>;
  getApiClient: () => Blog0ApiClient;
  fetchProfile: () => Promise<void>;
  // Efficient accessor methods
  isPostLiked: (postId: string) => boolean;
  isPostBookmarked: (postId: string) => boolean;
  isUserFollowed: (userId: string) => boolean;
  // Update methods
  updatePostLike: (postId: string, liked: boolean) => void;
  updatePostBookmark: (postId: string, bookmarked: boolean) => void;
  updateUserFollow: (userId: string, following: boolean) => void;
}

type AuthStore = AuthState & AuthActions;

export const useAuthStore = create<AuthStore>()(
  persist(
    (set, get) => ({
      // State
      user: null,
      token: null,
      isAuthenticated: false,
      isLoading: false,
      error: null,
      // Profile hash maps
      likedPosts: {},
      bookmarkedPosts: {},
      followingUsers: {},

      // Actions
      login: async (provider: string) => {
        try {
          set({ isLoading: true, error: null });
          const api = new Blog0ApiClient();

          // Start OAuth flow
          await api.startOAuth(provider);
        } catch (error) {
          set({
            error: error instanceof Error ? error.message : 'Login failed',
            isLoading: false,
          });
        }
      },

      logout: () => {
        // Clear store state
        set({
          user: null,
          token: null,
          isAuthenticated: false,
          error: null,
          isLoading: false,
          likedPosts: {},
          bookmarkedPosts: {},
          followingUsers: {},
        });
      },

      setToken: (token: string) => {
        set({
          token,
          isAuthenticated: true,
          error: null,
        });
        
        // Fetch profile data after setting token
        get().fetchProfile().catch(console.error);
      },

      setUser: (user: User) => {
        set({
          user,
          isAuthenticated: true,
          error: null,
        });
      },

      clearError: () => {
        set({ error: null });
      },

      setLoading: (loading: boolean) => {
        set({ isLoading: loading });
      },

      checkAuth: async () => {
        const { token } = get();

        if (!token) {
          set({ isAuthenticated: false, user: null });
          return;
        }

        try {
          set({ isLoading: true });

          // Try to fetch user data to verify token
          // For now, we'll just assume the token is valid
          set({
            isAuthenticated: true,
            isLoading: false,
            error: null,
          });
        } catch (error) {
          console.error('Auth check failed:', error);

          // Token is invalid, clear auth state
          set({
            user: null,
            token: null,
            isAuthenticated: false,
            isLoading: false,
            error: 'Authentication expired',
          });
        }
      },

      getApiClient: () => {
        const { token } = get();
        const api = new Blog0ApiClient();
        
        if (token) {
          api.setApiToken(token);
        }
        
        return api;
      },

      fetchProfile: async () => {
        const { token, isAuthenticated } = get();
        
        if (!token || !isAuthenticated) {
          return;
        }

        try {
          const api = get().getApiClient();
          const profile = await api.getProfile();
          
          // Convert arrays to hash maps for efficient lookups
          const likedPosts: Record<string, boolean> = {};
          const bookmarkedPosts: Record<string, boolean> = {};
          const followingUsers: Record<string, boolean> = {};
          
          profile.liked_posts.forEach(post => {
            likedPosts[post.id] = true;
          });
          
          profile.bookmarks.forEach(post => {
            bookmarkedPosts[post.id] = true;
          });
          
          profile.following.forEach(user => {
            followingUsers[user.id] = true;
          });
          
          set({ 
            likedPosts,
            bookmarkedPosts,
            followingUsers
          });
        } catch (error) {
          console.error('Failed to fetch profile:', error);
          // Don't clear auth state for profile fetch errors
        }
      },

      // Efficient accessor methods
      isPostLiked: (postId: string) => {
        const { likedPosts } = get();
        return likedPosts[postId] || false;
      },

      isPostBookmarked: (postId: string) => {
        const { bookmarkedPosts } = get();
        return bookmarkedPosts[postId] || false;
      },

      isUserFollowed: (userId: string) => {
        const { followingUsers } = get();
        return followingUsers[userId] || false;
      },

      // Update methods
      updatePostLike: (postId: string, liked: boolean) => {
        const { likedPosts } = get();
        const newLikedPosts = { ...likedPosts };
        
        if (liked) {
          newLikedPosts[postId] = true;
        } else {
          delete newLikedPosts[postId];
        }
        
        set({ likedPosts: newLikedPosts });
      },

      updatePostBookmark: (postId: string, bookmarked: boolean) => {
        const { bookmarkedPosts } = get();
        const newBookmarkedPosts = { ...bookmarkedPosts };
        
        if (bookmarked) {
          newBookmarkedPosts[postId] = true;
        } else {
          delete newBookmarkedPosts[postId];
        }
        
        set({ bookmarkedPosts: newBookmarkedPosts });
      },

      updateUserFollow: (userId: string, following: boolean) => {
        const { followingUsers } = get();
        const newFollowingUsers = { ...followingUsers };
        
        if (following) {
          newFollowingUsers[userId] = true;
        } else {
          delete newFollowingUsers[userId];
        }
        
        set({ followingUsers: newFollowingUsers });
      },
    }),
    {
      name: 'auth-storage',
      partialize: (state) => ({
        token: state.token,
        user: state.user,
        isAuthenticated: state.isAuthenticated,
        likedPosts: state.likedPosts,
        bookmarkedPosts: state.bookmarkedPosts,
        followingUsers: state.followingUsers,
      }),
    }
  )
);
