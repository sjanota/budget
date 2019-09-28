import React from 'react';
import ExpensesList from '../ExpensesList/ExpensesList';
import { Navigation } from '../Navigation/Navigation';
import './App.css';
import { Switch } from 'react-router-dom';
import { Route } from 'react-router-dom';
import { BudgetProvider } from '../context/budget/budget';
import { EnvelopesList } from '../EnvelopesList/EnvelopesList';
import { CategoriesList } from '../CategoriesList/CategoriesList';

export default function App() {
  return (
    <BudgetProvider>
      <div className="App">
        <Navigation />
        <div className={'App-main'}>
          <Switch>
            <Route path={'/expenses'} component={ExpensesList} />
            <Route
              path={'/accounts'}
              render={() => <p>Tu będzie lista kont</p>}
            />
            <Route path={'/envelopes'} component={EnvelopesList} />
            <Route path={'/categories'} component={CategoriesList} />
          </Switch>
        </div>
      </div>
    </BudgetProvider>
  );
}
