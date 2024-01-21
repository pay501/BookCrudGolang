/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    images:{
        remotePatterns:[
            {
                protocol:'https',
                hostname:'images.pexels.com'
            },
            {
                protocol:'https',
                hostname: "**"
            },
        ]
    }
}

module.exports = nextConfig
