import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import './css/MyTasks.css'

interface Todo {
  id: number;
  content: string;
}
const MyTasks: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [task, setTask] = useState<string>('');
  const navigate = useNavigate();

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then(response => response.json())
      .then(tasks => {
        if (tasks && Array.isArray(tasks)) {
          setTodos(tasks);
        } else {
          console.error('Unexpected response format:', tasks);
        }
      })
      .catch(error => console.error('Error fetching tasks:', error));
  }, []);



  const handleAddTasks = async () => {
    try {
      const response = await fetch(`http://localhost:8080/tasks`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: task })
      });

      if (!response.ok) {
        throw new Error('Failed to add task');
      }

      const result = await response.json();

      fetch('http://localhost:8080/tasks')
        .then(response => response.json())
        .then(tasks => setTodos(tasks))
        .catch(error => console.error('Error fetching tasks:', error));


      setTask('');
    } catch (error) {
      console.error('Error adding task:', error);
    }
  };
  const handleLogout = async () => {
    navigate('/');
  };
  return (
    <div className="mytasks">
      Task:
      <input
        type="text"
        value={task}
        onChange={(e) => setTask(e.target.value)}
        placeholder="Task" />
      <button onClick={handleAddTasks}>
        Add Task
      </button>
      Here are your Tasks:
      {todos.length === 0 ? (
        <p>No tasks set yet</p>
      ) : (<ul>
        {todos.map((todo) => (
          <li key={todo.id}>
            {todo.content}
          </li>
        ))}
      </ul>)}
      <button onClick={handleLogout}>
        Log Out
      </button>
    </div>
  );
};

export default MyTasks;