import React from 'react';
import { Route } from 'react-router-dom';
import Topbar from '../Topbar';
import SBAdmin2 from '../template/SBAdmin2';
import { sidebarConfig } from './sidebarConfig';
import { BudgetProvider, BudgetContext } from '../gql/BudgetContext';
import Accounts from '../Accounts';
import Envelopes from '../Envelopes/EnvelopesPage';
import Expenses from '../Expenses';

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
              <>
                <Route path="/accounts" component={Accounts} />
                <Route path="/envelopes" component={Envelopes} />
                <Route path="/expenses" component={Expenses} />
              </>
            )
          }
        </BudgetContext.Consumer>
      </SBAdmin2>
    </BudgetProvider>
  );
}
