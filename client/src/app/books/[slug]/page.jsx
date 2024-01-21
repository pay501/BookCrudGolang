"use client"

import axios from "axios"
import Image from "next/image"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"
import style from "./singleBook.module.css"

export default function singleBookPage ( {params} ) {

    const [book, setBook] = useState({})
    const router = useRouter()
    const id = params.slug

    async function getBook (){
        await axios.get(`http://localhost:8080/book/${id}`,{withCredentials:true})
        .then(res=>{
            if(res.data.message === "return"){
                router.push("/login")
            }else{
                setBook(res.data.result)
            }
        })
    }
    
    useEffect(()=>{
        getBook()
    },[])
    return (
        <div className={`${style.container} bg-base-300`}>
                <div className={`${style.imgContainer}`}>
                    <Image
                        src={book.image}
                        alt="Book"
                        fill
                        className={`${style.img}`}
                    >
                    </Image>
                </div>
                <div className={`${style.aboutContainer}`}>
                    <h1 className="text-4xl font-semibold text-white">Title - <span className="text-slate-300">{book.name}</span></h1>
                    <div className="ml-8 mt-4">
                        <p className="text-xl font-bold text-white">Author - <span className="text-slate-300">{book.author}</span></p>
                        <p className="text-lg font-semibold text-white mt-4">Description - <span className="text-slate-300">{book.description}</span></p>
                        <p className="text-lg font-semibold text-white mt-4">Added at - <span className="text-slate-300">{book && book.CreatedAt && book.CreatedAt.split("T")[0]}</span></p>
                    </div>
                </div>
        </div>
    )
}

//Lorem ipsum dolor, sit amet consectetur adipisicing elit. Iure, dignissimos. Rerum perferendis, delectus reprehenderit non at asperiores deserunt? Eligendi odio eos reiciendis unde sit aspernatur beatae est exercitationem perspiciatis culpa.