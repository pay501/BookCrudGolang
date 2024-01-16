"use client"

import axios from "axios"
import { useEffect, useState } from "react"

export default function singleBookPage ( {params} ) {

    const [book, setBook] = useState({})
    
    const id = params.slug

    async function getBook (){
        await axios.get(`http://localhost:8080/book/${id}`,{withCredentials:true})
        .then(res=>{
            setBook(res.data)
            console.log(res.data)
        })
    }


    useEffect(()=>{
        getBook()
    },[])
    return (
        <div>
            <div>
                <h1>{params.slug}</h1>
                <h1 className="text-4xl font-semibold text-white">{book.description}</h1>
            </div>
        </div>
    )
}