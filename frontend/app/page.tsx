'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { cookieUtils } from '@/lib/cookies';

export default function HomePage() {
  const router = useRouter();

  useEffect(() => {
    const userId = cookieUtils.getUserId();
    if (userId) {
      router.push('/tasks');
    } else {
      router.push('/login');
    }
  }, [router]);

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-xl">Cargando...</div>
    </div>
  );
}
