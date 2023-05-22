/** @type {import('next').NextConfig} */
const nextConfig = {
    async redirects() {
        return [
            {
                source: '/login',
                destination: process.env.LOGIN_DESTINATION,
                permanent: false,
                basePath: false
            },
        ]
    },
}

module.exports = nextConfig
