'use client';

import { useParams, useRouter, useSearchParams } from 'next/navigation';
import { useEffect, useState } from 'react';
import { useAuthStore } from '@/store/authStore';
import { useToast } from '@/components/ui/toast';

export default function AuthCallbackPage() {
  const params = useParams();
  const router = useRouter();
  const searchParams = useSearchParams();
  const { setToken, setUser, setLoading } = useAuthStore();
  const { showToast } = useToast();

  const [status, setStatus] = useState<'loading' | 'success' | 'error'>('loading');
  const [message, setMessage] = useState('Processing authentication...');

  useEffect(() => {
    const handleCallback = () => {
      try {
        const provider = params.provider as string;

        if (!provider) {
          throw new Error('Provider not specified');
        }

        setLoading(true);

        // Check URL parameters for token or error (backend format: ?token=%s&id=%s&email=%s&username=%s)
        const token = searchParams.get('token');
        const error = searchParams.get('error');
        const userId = searchParams.get('id');
        const userName = searchParams.get('username');
        const userEmail = searchParams.get('email');

        if (error) {
          throw new Error(error);
        }

        if (token) {
          // Set token in store
          setToken(token);

          // Set user if provided
          if (userId && userName) {
            setUser({
              id: userId,
              name: userName,
              email: userEmail || undefined,
            });
          }

          setStatus('success');
          setMessage('Authentication successful! Redirecting...');

          // Redirect to home page after a short delay
          setTimeout(() => {
            router.push('/');
          }, 2000);
        } else {
          throw new Error('No authentication token received');
        }
      } catch (error) {
        console.error('Auth callback error:', error);
        const errorMessage = error instanceof Error ? error.message : 'Authentication failed';
        setStatus('error');
        setMessage(errorMessage);
        showToast(errorMessage);

        // Redirect to home page after showing error
        setTimeout(() => {
          router.push('/');
        }, 3000);
      } finally {
        setLoading(false);
      }
    };

    handleCallback();
  }, [params.provider, searchParams, setToken, setUser, setLoading, router, showToast]);

  const getStatusColor = () => {
    switch (status) {
      case 'success':
        return 'text-[#25F4EE]';
      case 'error':
        return 'text-[#FE2C55]';
      default:
        return 'accent-text';
    }
  };

  const getStatusIcon = () => {
    switch (status) {
      case 'success':
        return '✓';
      case 'error':
        return '✗';
      default:
        return '⟳';
    }
  };

  return (
    <div className="post-container flex items-center justify-center mesh-background">
      <div className="text-center space-y-6 max-w-md px-6">
        <div className={`text-4xl ${getStatusIcon() === '⟳' ? 'animate-spin' : ''}`}>
          {getStatusIcon()}
        </div>

        <div className="space-y-2">
          <h1 className="text-2xl font-bold text-white">
            {status === 'success' && 'Welcome to Blog0!'}
            {status === 'error' && 'Authentication Failed'}
            {status === 'loading' && 'Authenticating...'}
          </h1>

          <p className={`${getStatusColor()} font-medium`}>{message}</p>
        </div>

        {status === 'loading' && (
          <div className="flex justify-center">
            <div className="animate-pulse flex space-x-1">
              <div className="h-2 w-2 bg-white rounded-full"></div>
              <div className="h-2 w-2 bg-white rounded-full animation-delay-200"></div>
              <div className="h-2 w-2 bg-white rounded-full animation-delay-400"></div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
