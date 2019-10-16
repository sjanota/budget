import React from 'react';
import Page from './template/Page/Page';
import { Gauge } from './template/Gauge';
import { useGetCurrentMonthlyReport } from './gql/monthlyReport';
import { WithQuery } from './gql/WithQuery';
import Amount from '../model/Amount';
import { Row } from 'react-bootstrap';
import { Panel } from './template/Utilities/Panel';

function Gauges({ className, month }) {
  return (
    <div className={className}>
      <Row>
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title="Planned budget"
          value={Amount.format(month.totalPlannedAmount)}
          faIcon="clipboard-list"
        />
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title="Incomes"
          value={Amount.format(month.totalIncomeAmount)}
          faIcon="briefcase"
        />
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title="Left to plan"
          value={Amount.format(
            month.totalIncomeAmount - month.totalPlannedAmount
          )}
          faIcon="balance-scale"
        />
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title="Expenses"
          value={Amount.format(month.totalExpenseAmount)}
          faIcon="receipt"
        />
      </Row>
    </div>
  );
}

function MonthSummary({ className }) {
  return (
    <Panel
      className={className}
      header={
        <div className="d-flex justify-content-between align-items-center">
          <Panel.Title>Summary</Panel.Title>
        </div>
      }
    />
  );
}

export function MonthDashboardPage() {
  const query = useGetCurrentMonthlyReport();
  return (
    <Page>
      <WithQuery query={query}>
        {({ data }) => (
          <Row>
            <Gauges
              className="col-12 col-lg-3"
              month={data.budget.currentMonth}
            />
            <MonthSummary className="col-12 col-lg-9 px-0" />
          </Row>
        )}
      </WithQuery>
    </Page>
  );
}
