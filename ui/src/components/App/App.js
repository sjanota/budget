import React, { useEffect } from 'react';
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
import { useAuth0 } from '../../react-auth0-spa';
import { LangProvider } from '../template/Utilities/Lang';
import pl from '../../lang/pl';

export default function App() {
  const { isAuthenticated, loginWithRedirect, loading } = useAuth0();

  useEffect(() => {
    if (loading) {
      return;
    }
    if (!isAuthenticated) {
      loginWithRedirect({});
      return;
    }
  }, [isAuthenticated, loginWithRedirect, loading]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated) {
    return <div />;
  }

  return (
    <LangProvider dictionary={pl}>
      <BudgetProvider>
        <SBAdmin2
          sidebarProps={{
            renderBrandName: () => 'Budget',
            renderBrandIcon: () => <i className="fas fa-bold" />,
            config: sidebarConfig(pl),
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
    </LangProvider>
  );
}
