import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import Blog0ApiClient from '@/lib/api-client';

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
        });
      },

      setToken: (token: string) => {
        set({
          token,
          isAuthenticated: true,
          error: null,
        });
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
    }),
    {
      name: 'auth-storage',
      partialize: (state) => ({
        token: state.token,
        user: state.user,
        isAuthenticated: state.isAuthenticated,
      }),
    }
  )
);
