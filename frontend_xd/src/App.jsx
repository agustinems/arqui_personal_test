import { useState } from 'react'
import './App.css'
import { Routes, Route } from 'react-router-dom'

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/admin" element={<Admin />} />
        <Route path="/user" element={<User />} />
      </Routes>
    </>
  )
}

export default App
