import React, { createContext, useState, useContext, useEffect } from 'react';
import PropTypes from 'prop-types';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

export const BudgetContext = createContext();
export const useBudget = () => useContext(BudgetContext);

const GET_BUDGETS = gql`
  query GetBudgets {
    budgets {
      id
      name
    }
  }
`;

const storageKey = 'LAST_CHOSEN-BUDGET-ID';

export function BudgetProvider({ children }) {
  const [selectedBudget, setSelectedBudget] = useState(null);
  const { loading, error, data } = useQuery(GET_BUDGETS, {
    pollInterval: 10000,
  });
  useEffect(() => {
    if (selectedBudget) {
      sessionStorage.setItem(storageKey, selectedBudget.id);
    }
  }, [selectedBudget]);
  useEffect(() => {
    if (!selectedBudget && data && data.budgets) {
      const lastChosenID = sessionStorage.getItem(storageKey);
      const lastChosen = data.budgets.find(b => b.id === lastChosenID);
      if (lastChosen) {
        setSelectedBudget(lastChosen);
      }
    }
  }, [data, selectedBudget]);
  const value = {
    selectedBudget,
    setSelectedBudget,
    loading,
    error,
    budgets: !loading && !error ? data.budgets : [],
  };
  if (error) {
    console.error(error);
  }
  return (
    <BudgetContext.Provider value={value}>{children}</BudgetContext.Provider>
  );
}

BudgetProvider.propTypes = {
  children: PropTypes.node,
};
