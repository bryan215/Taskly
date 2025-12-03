'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { cookieUtils } from '@/lib/cookies';

export default function HomePage() {
  const router = useRouter();

  useEffect(() => {
    const token = cookieUtils.getToken();
    
    // Verificar que exista token y no esté expirado
    if (token && !cookieUtils.isTokenExpired()) {
      router.push('/tasks');
    } else {
      if (token) {
        cookieUtils.clearAuth(); // Limpiar token expirado
      }
      router.push('/login');
    }
  }, [router]);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="text-center">
        <h1 className="text-5xl font-bold text-indigo-600 mb-2">Taskly</h1>
        <p className="text-gray-600">Cargando...</p>
      </div>
    </div>
  );
}
