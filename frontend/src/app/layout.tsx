

'use client';
import 'bootstrap/dist/css/bootstrap.min.css';
import { AuthProvider } from "@/modules/auth/core/context"

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <AuthProvider>
          {children}
        </AuthProvider>
      </body>
    </html>
  )
}
