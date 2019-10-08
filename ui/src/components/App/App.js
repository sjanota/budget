import React from 'react';
import Dashboard from '../Dashboard';
import Buttons from '../Buttons';
import Tables from '../Tables';
import { Route } from 'react-router-dom';
import Topbar from '../Topbar';
import SBAdmin2 from '../template/SBAdmin2';
import { sidebarConfig } from './sidebarConfig';
import { BudgetProvider, BudgetContext } from '../contexts/BudgetContext';
import Accounts from '../Accounts';

export default function App() {
  return (
    <BudgetProvider>
      <SBAdmin2
        sidebarProps={{
          renderBrandName: () => (
            <>
              SB Admin <sup>2</sup>
            </>
          ),
          renderBrandIcon: () => <i className="fas fa-laugh-wink" />,
          config: sidebarConfig,
        }}
        topbar={Topbar}
        copyright={'Budget 2019'}
      >
        <BudgetContext.Consumer>
          {({ selectedBudget }) =>
            selectedBudget && (
              <>
                <Route path="/buttons" component={Buttons} />
                <Route path="/tables" component={Tables} />
                <Route path="/accounts" component={Accounts} />
                <Route exact path="/" component={Dashboard} />
              </>
            )
          }
        </BudgetContext.Consumer>
      </SBAdmin2>
    </BudgetProvider>
  );
}
