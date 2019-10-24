import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Topbar from '../Topbar';
import SBAdmin2 from '../template/SBAdmin2';
import { sidebarConfig } from './sidebarConfig';
import { BudgetProvider, BudgetContext } from '../gql/budget';
import Accounts from '../Accounts';
import Envelopes from '../Envelopes/EnvelopesPage';
import Expenses from '../Expenses';
import Transfers from '../Transfers';
import Plans from '../Plans';
import { MonthDashboardPage } from '../MonthDashboardPage';

export default function App() {
  return (
    <BudgetProvider>
      <SBAdmin2
        sidebarProps={{
          renderBrandName: () => 'Budget',
          renderBrandIcon: () => <i className="fas fa-bold" />,
          config: sidebarConfig,
        }}
        topbar={Topbar}
        copyright={'Budget 2019'}
      >
        <BudgetContext.Consumer>
          {({ selectedBudget }) =>
            selectedBudget && (
              <Switch>
                <Route path="/accounts" component={Accounts} />
                <Route path="/envelopes" component={Envelopes} />
                <Route path="/expenses" component={Expenses} />
                <Route path="/transfers" component={Transfers} />
                <Route path="/plans" component={Plans} />
                <Route path="/" component={MonthDashboardPage} />
              </Switch>
            )
          }
        </BudgetContext.Consumer>
      </SBAdmin2>
    </BudgetProvider>
  );
}
