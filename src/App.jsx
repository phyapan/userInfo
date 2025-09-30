
import { useState } from 'react';
import './App.css';

function App() {
  const [form, setForm] = useState({
    name: '',
    age: '',
    email: '',
    contact: ''
  });
  const [submitted, setSubmitted] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/submit', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams(form).toString(),
      });
      if (response.ok) {
        setSubmitted(true);
      } else {
        alert('Failed to submit data');
      }
    } catch (error) {
      alert('Error submitting data');
    }
  };

  return (
    <div className="form-container">
      <h1>User Information Form</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name:</label>
          <input type="text" name="name" value={form.name} onChange={handleChange} required />
        </div>
        <div>
          <label>Age:</label>
          <input type="number" name="age" value={form.age} onChange={handleChange} required />
        </div>
        <div>
          <label>Email:</label>
          <input type="email" name="email" value={form.email} onChange={handleChange} required />
        </div>
        <div>
          <label>Contact:</label>
          <input type="text" name="contact" value={form.contact} onChange={handleChange} required />
        </div>
        <button type="submit">Submit</button>
      </form>
      {/* Submitted data display removed as requested */}
    </div>
  );
}

export default App;
