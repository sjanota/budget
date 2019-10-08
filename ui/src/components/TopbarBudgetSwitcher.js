import React from 'react';
import TopbarContextSwitcher from './template/Topbar/TopbarContextSwitcher';
import { useBudget } from './contexts/BudgetContext';
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
    selectedBudget
  );
  return (
    <TopbarContextSwitcher
      label="Budget"
      value={value}
      onChange={setSelectedBudget}
      allowedValues={budgets.map(b => b.name)}
      loadingValues={loading || !!error}
    />
  );
}
