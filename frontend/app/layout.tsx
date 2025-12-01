import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Taskly - Gestor de Tareas",
  description: "Aplicación moderna para gestionar tus tareas",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="es">
      <body className="antialiased">
        {children}
      </body>
    </html>
  );
}
