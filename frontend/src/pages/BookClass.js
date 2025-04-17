import React, { useState } from "react";
import axios from "axios";

function BookClass() {
  const [user, setUser] = useState("");
  const [classId, setClassId] = useState("");

  const book = () => {
    axios.post("http://localhost:5000/api/book", { user, classId })
      .then((res) => alert(res.data.message))
      .catch(console.error);
  };

  return (
    <div className="p-6">
      <h2 className="text-xl mb-2">Book a Class</h2>
      <input
        placeholder="Your name"
        className="border p-2 mr-2"
        value={user}
        onChange={(e) => setUser(e.target.value)}
      />
      <input
        placeholder="Class ID"
        className="border p-2 mr-2"
        value={classId}
        onChange={(e) => setClassId(e.target.value)}
      />
      <button onClick={book} className="bg-blue-500 text-white p-2">Book</button>
    </div>
  );
}

export default BookClass;
