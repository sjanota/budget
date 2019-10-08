import React, { useState } from 'react';
import TopbarContextSwitcher from './template/Topbar/TopbarContextSwitcher';

const budgets = ['G.I Joe - Personal', 'US Army'];

export default function TopbarBudgetSwitcher() {
  const [budget, setBudget] = useState(null);
  return (
    <TopbarContextSwitcher
      label="Budget"
      value={budget}
      onChange={setBudget}
      allowedValues={budgets}
    />
  );
}
