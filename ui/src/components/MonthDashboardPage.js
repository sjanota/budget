import React from 'react';
import Page from './template/Page/Page';
import { Gauge } from './template/Gauge';
import { useGetCurrentMonthlyReport } from './gql/monthlyReport';
import { WithQuery } from './gql/WithQuery';
import { useGetEnvelopes } from './gql/envelopes';
import { useGetAccounts } from './gql/accounts';
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

const severityIcon = {
  ERROR: 'exclamation-circle',
  WARNING: 'exclamation-triangle',
};

const severityVariant = {
  ERROR: 'danger',
  WARNING: 'warning',
};

function ProblemMessage({ problem }) {
  const envelopesQuery = useGetEnvelopes();
  const accountsQuery = useGetAccounts();

  return (
    <WithQuery query={envelopesQuery}>
      {({ data: envelopesData }) => (
        <WithQuery query={accountsQuery}>
          {({ data: accountsData }) =>
            problem.__typename === 'Misplanned'
              ? problem.overplanned
                ? 'Plans for this month exceeds incomes'
                : 'There are unplanned resources'
              : problem.__typename === 'NegativeBalanceOnEnvelope'
              ? `Expenses have exceeded plans for envelope "${
                  envelopesData.envelopes.find(e => e.id === problem.id).name
                }"`
              : problem.__typename === 'EnvelopeOverLimit'
              ? `Plans for envelope "${
                  envelopesData.envelopes.find(e => e.id === problem.id).name
                }" exceed the limit`
              : problem.__typename === 'NegativeBalanceOnAccount'
              ? `Expenses have exceeded incomes on account "${
                  accountsData.accounts.find(a => a.id === problem.id).name
                }"`
              : problem.__typename === 'MonthStillInProgress'
              ? 'Month has not ended yet'
              : problem.__typename
          }
        </WithQuery>
      )}
    </WithQuery>
  );
}

function MonthSummary({ className, problems }) {
  return (
    <Panel
      className={className}
      header={
        <div className="d-flex justify-content-between align-items-center">
          <Panel.Title>Problems</Panel.Title>
        </div>
      }
      body={
        <ul className="list-group list-group-flush">
          {problems.length > 0 ? (
            problems.map((problem, idx) => (
              <li
                key={idx}
                className={`list-group-item text-${
                  severityVariant[problem.severity]
                }`}
              >
                <i
                  className={`fas fa-fw fa-${
                    severityIcon[problem.severity]
                  } mr-1`}
                />
                <ProblemMessage problem={problem} />
              </li>
            ))
          ) : (
            <li className="list-group-item text-success">
              <i className="fas fa-fw fa-check-circle mr-1" />
              Everything is fine
            </li>
          )}
        </ul>
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
            <MonthSummary
              className="col-12 col-lg-9 px-0"
              problems={data.budget.currentMonth.problems}
            />
          </Row>
        )}
      </WithQuery>
    </Page>
  );
}
