import React from 'react';
import Page from './template/Page/Page';
import { Gauge } from './template/Gauge';
import { useGetCurrentMonthlyReport } from './gql/monthlyReport';
import { WithQuery } from './gql/WithQuery';
import Amount from '../model/Amount';
import { Row } from 'react-bootstrap';

export function MonthDashboardPage() {
  const query = useGetCurrentMonthlyReport();
  return (
    <Page>
      <WithQuery query={query}>
        {({ data }) => (
          <Row>
            <Gauge
              className="col-xl-3 col-md-6 mb-4"
              variant="primary"
              title="Planned budget"
              value={Amount.format(data.budget.currentMonth.totalPlannedAmount)}
              faIcon="receipt"
            />
            <Gauge
              className="col-xl-3 col-md-6 mb-4"
              variant="primary"
              title="Incomes"
              value={Amount.format(data.budget.currentMonth.totalIncomeAmount)}
              faIcon="receipt"
            />
            <Gauge
              className="col-xl-3 col-md-6 mb-4"
              variant="primary"
              title="Left to plan"
              value={Amount.format(
                data.budget.currentMonth.totalIncomeAmount -
                  data.budget.currentMonth.totalPlannedAmount
              )}
              faIcon="receipt"
            />
          </Row>
        )}
      </WithQuery>
    </Page>
  );
}
