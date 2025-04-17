import React, { useEffect, useState } from "react";
import axios from "axios";

function Classes() {
  const [classes, setClasses] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:5000/api/classes")
      .then(res => setClasses(res.data))
      .catch(console.error);
  }, []);

  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Available Classes</h2>
      <ul>
        {classes.map((c) => (
          <li key={c.id} className="mb-2 p-4 border rounded">
            <strong>{c.name}</strong> by {c.instructor} - ${c.price}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Classes;
