'use client';

import { Mail, X } from 'lucide-react';
import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { useAuthStore } from '@/store/authStore';

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export default function LoginModal({ isOpen, onClose }: LoginModalProps) {
  const { login, isLoading, error, clearError } = useAuthStore();
  const [selectedProvider, setSelectedProvider] = useState<string | null>(null);

  if (!isOpen) return null;

  const handleLogin = async (provider: string) => {
    setSelectedProvider(provider);
    clearError();

    try {
      await login(provider);
      // OAuth will redirect, so we don't close the modal here
    } catch (err) {
      console.error('Login failed:', err);
      setSelectedProvider(null);
    }
  };

  const handleClose = () => {
    if (!isLoading) {
      clearError();
      setSelectedProvider(null);
      onClose();
    }
  };

  return (
    <div className="fixed inset-0 bg-black/80 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <Card className="w-full max-w-md bg-[#121212] border-white/10">
        <CardHeader className="text-center relative">
          <Button
            variant="ghost"
            size="sm"
            className="absolute right-2 top-2 text-[#AFAFAF] hover:text-white hover:bg-white/10"
            onClick={handleClose}
            disabled={isLoading}
          >
            <X className="h-4 w-4" />
          </Button>
          <CardTitle className="text-white text-xl font-bold mb-2">Join Blog0</CardTitle>
          <p className="text-[#AFAFAF] text-sm">Sign in to like, comment, and share posts</p>
        </CardHeader>

        <CardContent className="space-y-6">
          {error && (
            <div className="bg-[#FE2C55]/10 border border-[#FE2C55]/20 rounded p-3">
              <p className="text-[#FE2C55] text-sm">{error}</p>
            </div>
          )}

          <div className="flex justify-center">
            <Button
              onClick={() => handleLogin('google')}
              disabled={isLoading}
              className={`w-full max-w-sm bg-white text-black border-0 hover:bg-gray-100 transition-smooth ${
                selectedProvider === 'google' && isLoading ? 'opacity-50' : ''
              }`}
            >
              <Mail className="h-4 w-4 mr-2" />
              {selectedProvider === 'google' && isLoading
                ? 'Connecting...'
                : 'Continue with Google'}
            </Button>
          </div>

          <div className="text-center pt-4">
            <p className="text-[#AFAFAF] text-xs">
              By continuing, you agree to our Terms of Service and Privacy Policy
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
