import React from 'react';
import Dashboard from '../Dashboard';
import Buttons from '../Buttons';
import Tables from '../Tables';
import { Route } from 'react-router-dom';
import Topbar from '../Topbar';
import SBAdmin2 from '../template/SBAdmin2';
import { sidebarConfig } from './sidebarConfig';

export default function App() {
  return (
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
      <Route path="/buttons" component={Buttons} />
      <Route path="/tables" component={Tables} />
      <Route exact path="/" component={Dashboard} />
    </SBAdmin2>
  );
}
