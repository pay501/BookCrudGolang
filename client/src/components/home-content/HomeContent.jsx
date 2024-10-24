"use client"
import axios from "axios"
import Image from "next/image"
import Link from "next/link"
import { useEffect, useState } from "react"
import style from "./homecontent.module.css"

export default function HomeContent() {

    const [items, setItems] = useState([])
    const [isLoading, setIsLoading] = useState(false)
    
    async function getBooks() {
        try {
            const res = await axios.get("http://localhost:8080/books", { withCredentials: true });
            setItems(res.data.result);
        } catch (error) {
            console.error("Error fetching books:", error);
        }
        setIsLoading(false)
    }
    useEffect(() => {
        setIsLoading(true)
        getBooks()
    }, [])


    return (
        <div>
            {
                isLoading ?
                <div className="bg-black h-screen opacity-50 flex flex-col justify-center items-center">
                    <span className="loading loading-spinner loading-lg"></span>
                    <div className="text-xl font-semibold">Loading...</div>
                </div>
                :
                <div className={`${style.container}`}>
                <section className="hero">
                    <div className="hero bg-base-200 py-10">
                        <div className="hero-content text-center">
                            <div className="max-w-md">
                                <h1 className="text-5xl font-bold">Hello there</h1>
                                <p className="py-6">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</p>
                                <button className="btn btn-primary">Get Started</button>
                            </div>
                        </div>
                    </div>
                </section>
                {
                    items.length > 0 ? 
                    <section className={`${style.card} `}>
                    {
                        items.map((val, key) => {
                            return (
                                <div className={style.insideCard} key={key}>
                                    <div className="card card-side bg-base-100 shadow-xl w-full">
                                        <div className={style.imgContainer}>
                                            <Image
                                                src={val.image}
                                                alt="Movie"
                                                fill
                                                className={`${style.img}`}
                                            >
                                            </Image>
                                        </div>
                                        <div className={`${style.textContainer}`}>
                                            <div className="card-body">
                                                <h2 className="card-title text-2xl font-bold text-white">{val.name}</h2>
                                                <p className={`font-semibold`}>Written by {val.author}</p>
                                                <div className="card-actions justify-end">
                                                    <Link className="btn btn-primary" href={`/books/${val.ID}`}>More</Link>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            )
                        })
                    }
                </section>
                :
                <section></section>
                }
            </div>
            }
        </div>
    )
}


{/* <div className="card w-96 bg-base-100 shadow-xl">
    <figure className="px-10 pt-10">
        <img src="https://daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg" alt="Shoes" className="rounded-xl" />
    </figure>
    <div className="card-body items-center text-center">
        <h2 className="card-title">Shoes!</h2>
        <p>If a dog chews shoes whose shoes does he choose?</p>
        <div className="card-actions">
            <button className="btn btn-primary">Buy Now</button>
        </div>
    </div>
</div> */}