import React from 'react';
import TopbarContextSwitcher from './template/Topbar/TopbarContextSwitcher';
import { useBudget } from './gql/BudgetContext';
import Spinner from './template/Utilities/Spinner';

export default function TopbarBudgetSwitcher() {
  const {
    selectedBudget,
    setSelectedBudget,
    budgets,
    loading,
    error,
  } = useBudget();
  const value = loading ? (
    <Spinner size="sm" variant="secondary" />
  ) : error ? (
    <i className="fas fa-fw fa-exclamation-triangle text-secondary" />
  ) : (
    selectedBudget && selectedBudget.name
  );
  const onChange = id => {
    const budget = budgets.find(b => b.id === id);
    setSelectedBudget(budget);
  };
  return (
    <TopbarContextSwitcher
      label="Budget"
      value={value}
      onChange={onChange}
      allowedValues={budgets.map(b => ({ id: b.id, label: b.name }))}
    />
  );
}
