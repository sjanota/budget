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
import { SplitButton } from './template/Utilities/SplitButton';
import Month from '../model/Month';
import { useCloseCurrentMonth } from './gql/budget';
import { useDictionary } from './template/Utilities/Lang';

function Gauges({ className, month }) {
  const { dashboard } = useDictionary();
  return (
    <div className={className}>
      <Row>
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title={dashboard.planned}
          value={Amount.format(month.totalPlannedAmount)}
          faIcon="clipboard-list"
        />
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title={dashboard.incomes}
          value={Amount.format(month.totalIncomeAmount)}
          faIcon="briefcase"
        />
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title={dashboard.leftToPlan}
          value={Amount.format(
            month.totalIncomeAmount - month.totalPlannedAmount
          )}
          faIcon="balance-scale"
        />
        <Gauge
          className="col-6 col-lg-12 mb-4"
          variant="primary"
          title={dashboard.expenses}
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
  INFO: 'info-circle',
};

const severityVariant = {
  ERROR: 'danger',
  WARNING: 'warning',
  INFO: 'primary',
};

function ProblemMessage({ problem }) {
  const envelopesQuery = useGetEnvelopes();
  const accountsQuery = useGetAccounts();
  const { dashboard } = useDictionary();

  return (
    <WithQuery query={envelopesQuery}>
      {({ data: envelopesData }) => (
        <WithQuery query={accountsQuery}>
          {({ data: accountsData }) =>
            problem.__typename === 'Misplanned'
              ? problem.overplanned
                ? dashboard.problems.overplanned
                : dashboard.problems.underplanned
              : problem.__typename === 'NegativeBalanceOnEnvelope'
              ? dashboard.problems.expensesExceedPlans(
                  envelopesData.envelopes.find(e => e.id === problem.id).name
                )
              : problem.__typename === 'EnvelopeOverLimit'
              ? dashboard.problems.envelopeOverLimit(
                  envelopesData.envelopes.find(e => e.id === problem.id).name
                )
              : problem.__typename === 'NegativeBalanceOnAccount'
              ? dashboard.problems.negativeAccountBalance(
                  accountsData.accounts.find(a => a.id === problem.id).name
                )
              : problem.__typename === 'MonthStillInProgress'
              ? dashboard.problems.monthNotEnded
              : problem.__typename
          }
        </WithQuery>
      )}
    </WithQuery>
  );
}

function Problem({ problem }) {
  return (
    <li className={`list-group-item text-${severityVariant[problem.severity]}`}>
      <i className={`fas fa-fw fa-${severityIcon[problem.severity]} mr-1`} />
      <ProblemMessage problem={problem} />
    </li>
  );
}

function NoProblems() {
  const { dashboard } = useDictionary();
  return (
    <li className="list-group-item text-success">
      <i className="fas fa-fw fa-check-circle mr-1" />
      {dashboard.noProblems}
    </li>
  );
}

function MonthProblems({ className, problems }) {
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
              <Problem key={idx} problem={problem} />
            ))
          ) : (
            <NoProblems />
          )}
        </ul>
      }
    />
  );
}

function StartNextMonthButton({ disabled, warn }) {
  const [closeCurrentMonth] = useCloseCurrentMonth();
  const { dashboard } = useDictionary();
  return (
    <SplitButton
      faIcon="clipboard-check"
      variant={warn ? 'warning' : 'success'}
      disabled={disabled}
      onClick={() => closeCurrentMonth()}
    >
      {dashboard.buttons.closeMonth}
    </SplitButton>
  );
}

function CurrentMonth({ className, month }) {
  const { dashboard, months } = useDictionary();
  const parsed = Month.parse(month.month);

  return (
    <Panel
      className={className}
      header={
        <Panel.HeaderWithButton
          title={
            <span>
              {dashboard.currentMonth}:{' '}
              <strong>
                <em>
                  {months[parsed.month - 1]} {parsed.year}
                </em>
              </strong>
            </span>
          }
        >
          <StartNextMonthButton
            disabled={month.problems.some(p => p.severity === 'ERROR')}
            warn={month.problems.length > 0}
          />
        </Panel.HeaderWithButton>
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
            <CurrentMonth
              className="col-12 d-lg-none px-0"
              month={data.budget.currentMonth}
            />
            <Gauges
              className="col-12 col-lg-3"
              month={data.budget.currentMonth}
            />
            <MonthProblems
              className="col-12 d-lg-none px-0"
              problems={data.budget.currentMonth.problems}
            />
            <Row className="col-12 col-lg-9 flex-lg-column">
              <CurrentMonth
                className="d-none d-lg-block"
                month={data.budget.currentMonth}
              />
              <MonthProblems
                className="d-none d-lg-block flex-grow-1"
                problems={data.budget.currentMonth.problems}
              />
            </Row>
          </Row>
        )}
      </WithQuery>
    </Page>
  );
}
