import React from 'react';
import ExpensesList from '../ExpensesList/ExpensesList';
import { Navigation } from '../Navigation/Navigation';
import './App.css';
import { Switch } from 'react-router-dom';
import Route from 'react-router-dom/es/Route';

export default function App() {
  return (
    <div className="App">
      <Navigation />
      <div className={'App-main'}>
        <Switch>
          <Route path={'/expenses'} component={ExpensesList} />
        </Switch>
      </div>
    </div>
  );
}
