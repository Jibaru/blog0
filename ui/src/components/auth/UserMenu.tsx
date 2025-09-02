'use client';

import { LogOut, Settings, User } from 'lucide-react';
import { useState } from 'react';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { useAuthStore } from '@/store/authStore';

export default function UserMenu() {
  const { user, logout } = useAuthStore();
  const [isOpen, setIsOpen] = useState(false);

  if (!user) return null;

  const handleLogout = () => {
    logout();
    setIsOpen(false);
  };

  return (
    <div className="relative">
      <Button variant="ghost" size="sm" className="p-1" onClick={() => setIsOpen(!isOpen)}>
        <Avatar className="w-8 h-8 border border-white/20">
          <AvatarFallback className="bg-[#FE2C55] text-white font-bold text-sm">
            {user.name.charAt(0).toUpperCase()}
          </AvatarFallback>
        </Avatar>
      </Button>

      {isOpen && (
        <>
          {/* Backdrop */}
          <div className="fixed inset-0 z-40" onClick={() => setIsOpen(false)} />

          {/* Menu */}
          <Card className="absolute top-full right-0 mt-2 w-48 bg-[#121212] border-white/10 z-50">
            <CardContent className="p-2">
              <div className="flex items-center space-x-3 p-2 mb-2">
                <Avatar className="w-10 h-10 border border-white/20">
                  <AvatarFallback className="bg-[#FE2C55] text-white font-bold">
                    {user.name.charAt(0).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div>
                  <div className="text-white font-medium text-sm">{user.name}</div>
                  {user.email && <div className="text-[#AFAFAF] text-xs">{user.email}</div>}
                </div>
              </div>

              <div className="space-y-1">
                <Button
                  variant="ghost"
                  size="sm"
                  className="w-full justify-start text-[#AFAFAF] hover:text-white hover:bg-white/10"
                >
                  <User className="h-4 w-4 mr-2" />
                  Profile
                </Button>

                <Button
                  variant="ghost"
                  size="sm"
                  className="w-full justify-start text-[#AFAFAF] hover:text-white hover:bg-white/10"
                >
                  <Settings className="h-4 w-4 mr-2" />
                  Settings
                </Button>

                <div className="border-t border-white/10 my-1" />

                <Button
                  variant="ghost"
                  size="sm"
                  className="w-full justify-start text-[#FE2C55] hover:text-[#FE2C55] hover:bg-[#FE2C55]/10"
                  onClick={handleLogout}
                >
                  <LogOut className="h-4 w-4 mr-2" />
                  Logout
                </Button>
              </div>
            </CardContent>
          </Card>
        </>
      )}
    </div>
  );
}
