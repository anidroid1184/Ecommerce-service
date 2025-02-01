import React from 'react';
import ProductList from './components/ProductList';
import AddProduct from './components/AddProduct';

function App() {
  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <header className="text-center mb-8">
        <h1 className="text-4xl font-bold text-gray-800">Ecommerce App</h1>
        <p className="text-gray-600 mt-2">Gestiona tus productos fácilmente</p>
      </header>
      <main className="max-w-4xl mx-auto bg-white p-6 rounded-lg shadow-md">
        <AddProduct />
        <hr className="my-6 border-t border-gray-300" />
        <ProductList />
      </main>
    </div>
  );
}

export default App;