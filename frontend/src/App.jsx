import React from 'react';
import ProductList from './components/ProductList';
import AddProduct from './components/AddProduct';

function App() {
  return (
    <div>
      <h1>Ecommerce App</h1>
      <AddProduct />
      <ProductList />
    </div>
  );
}

export default App;
