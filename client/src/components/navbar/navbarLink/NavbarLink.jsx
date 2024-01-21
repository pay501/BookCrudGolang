"use client"

import axios from "axios"
import Link from "next/link"
import { usePathname, useRouter } from "next/navigation"
import { useEffect, useState } from "react"
import style from "./navbarLink.module.css"

export default function NavbarLinkPage() {

    const pathName = usePathname()
    const link = [
        {
            name: "home",
            path: "/"
        },
        {
            name: "management",
            path:"/manage/book"
        },
        {
            name: "addBook",
            path: "/books/addBook"
        },
    ]

    const [isSession, setIsSession] = useState(false)

    const router = useRouter()

    async function checkLogin () {
        await axios.get("http://localhost:8080/checkSession",{withCredentials:true})
        .then(res=>{
            if(res.data.message === "Session true"){
                setIsSession(true)
            }else{
                setIsSession(false)
            }
        })
    }

    async function logout () {
        await axios.get("http://localhost:8080/logout",{withCredentials:true})
        .then(res=>{
            if(res.data.message === "Logout successful"){
                setIsSession(false)
                router.push("/")
            }
        })
    }
    checkLogin()
    useEffect(()=>{
        checkLogin()
    },[])

    return (
        <div className={`${style.container} flex gap-2 justify-center items-center`}>
            <div className="flex gap-4">
                {
                    link.map((val, key) => {
                        return (
                            <div key={key}>
                                <Link
                                    href={val.path}
                                    className={`
                                        text-md font-semibold w-24 p-2 rounded-lg hover:text-blue-500
                                        ${pathName === val.path ? "bg-blue-500 text-black hover:text-white" : ""}
                                    `}
                                >
                                    {val.name}
                                </Link>
                            </div>
                        )
                    })
                }
            </div>
            <div>
                {
                    isSession ?
                    <button
                        onClick={logout}
                        className={`
                            text-md font-semibold w-20 rounded-lg hover:text-red-500
                        `}
                >
                    Logout
                </button>
                    :
                    <Link
                        href="/login"
                        className={`
                                text-md font-semibold w-24 p-2 rounded-lg hover:text-blue-500
                                ${pathName === "/login" ?"bg-blue-500 text-black hover:text-white" : "none"}
                        `}
                    >
                        Login
                    </Link>
                }
            </div>
        </div>
    )
}