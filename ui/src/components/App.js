import React from 'react';
import Dashboard from './Dashboard';
import Buttons from './Buttons';
import { Switch, Route } from 'react-router-dom';
import Sidebar from './Sidebar';
import Navbar from './Navbar';
import SBAdmin2 from './template/SBAdmin2';

export default function App() {
  return (
    <SBAdmin2>
      <Sidebar />
      <div id="content-wrapper" className="d-flex flex-column">
        <div id="content">
          <Navbar />
          <Switch>
            <Route path="/buttons" component={Buttons} />
            <Route path="/" component={Dashboard} />
          </Switch>
        </div>

        <footer className="sticky-footer bg-white">
          <div className="container my-auto">
            <div className="copyright text-center my-auto">
              <span>Copyright &copy; Your Website 2019</span>
            </div>
          </div>
        </footer>
      </div>
    </SBAdmin2>
  );
}
