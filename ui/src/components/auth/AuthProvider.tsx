'use client';

import { useEffect, useRef } from 'react';
import { useAuthStore } from '@/store/authStore';

interface AuthProviderProps {
  children: React.ReactNode;
}

export default function AuthProvider({ children }: AuthProviderProps) {
  const { checkAuth, checkTokenExpiration, isAuthenticated } = useAuthStore();
  const intervalRef = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
    // Check authentication status on app load
    checkAuth();
  }, [checkAuth]);

  // Set up periodic token expiration checks
  useEffect(() => {
    if (isAuthenticated) {
      // Check token expiration every 30 seconds
      intervalRef.current = setInterval(() => {
        checkTokenExpiration();
      }, 30000);
    } else {
      // Clear interval when not authenticated
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
        intervalRef.current = null;
      }
    }

    // Cleanup interval on unmount
    return () => {
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
      }
    };
  }, [isAuthenticated, checkTokenExpiration]);

  // Also check token expiration when the page gains focus
  useEffect(() => {
    const handleFocus = () => {
      if (isAuthenticated) {
        checkTokenExpiration();
      }
    };

    window.addEventListener('focus', handleFocus);
    return () => window.removeEventListener('focus', handleFocus);
  }, [isAuthenticated, checkTokenExpiration]);

  return <>{children}</>;
}
