import { NextResponse } from "next/server";

export async function middleware (request) {
    
    try{
        const jwt = request.cookies.get('jwt').value
        if(!jwt){
            return NextResponse.redirect(new URL('/login', request.url))
        }
        const requestHeaders = new Headers(request.headers)
        requestHeaders.set('user', JSON.stringify({ jwt:jwt }))
        const response = NextResponse.next({
            request:{
                headers:requestHeaders
            },
        })
        console.log(jwt)
        return response
    }catch(err){
        return NextResponse.redirect(new URL('/login', request.url))
    }
}

export const config = {
    matcher:[
        '/books/:path*',
    ]
}