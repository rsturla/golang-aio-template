// @ts-check

const isDevelopment = process.env.NODE_ENV === 'development';

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  output: isDevelopment ? 'standalone' : 'export',
  distDir: 'dist',
  // Use a function to conditionally set rewrites only in development
  ...(isDevelopment && {
    async rewrites() {
      return [
        {
          source: '/api/:path*',
          destination: 'http://localhost:8080/api/:path*',
        },
      ];
    },
  }),
};

module.exports = nextConfig;
