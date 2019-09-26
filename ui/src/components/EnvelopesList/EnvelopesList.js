import React from 'react';
import List from '../common/List/List';
import { useBudget } from '../context/budget/budget';
import { useQuery, useMutation } from '@apollo/react-hooks';
import {
  QUERY_ENVELOPES,
  UPDATE_ENVELOPE,
  CREATE_ENVELOPE,
} from './EnvelopesList.gql';
import { Envelope } from '../../model/propTypes';
import PropTypes from 'prop-types';

export function EnvelopesList() {
  const { id: budgetID } = useBudget();
  const { loading, error, data } = useQuery(QUERY_ENVELOPES, {
    variables: { budgetID },
  });
  const [updateEnvelope] = useMutation(UPDATE_ENVELOPE, {
    refetchQueries: () => [
      {
        query: QUERY_ENVELOPES,
        variables: { budgetID },
      },
    ],
  });
  const [createEnvelope] = useMutation(CREATE_ENVELOPE, {
    refetchQueries: () => [
      {
        query: QUERY_ENVELOPES,
        variables: { budgetID },
      },
    ],
  });

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

  return (
    <div>
      <List
        emptyValue={{ name: '' }}
        entries={data.envelopes}
        onCreate={input =>
          createEnvelope({
            variables: { budgetID, input: prepareInput(input) },
          })
        }
        onUpdate={(id, input) =>
          updateEnvelope({
            variables: { budgetID, id, input: prepareInput(input) },
          })
        }
        renderHeader={() => <ListHeader />}
        renderEntry={props => <ListEntry {...props} />}
        renderEditEntry={props => <EditEntry {...props} />}
      />
    </div>
  );
}

function ListHeader() {
  return (
    <>
      <th>Nazwa</th>
    </>
  );
}

function ListEntry({ entry }) {
  return (
    <>
      <td>{entry.name}</td>
    </>
  );
}

ListEntry.propTypes = {
  entry: Envelope.isRequired,
};

function EditEntry({ entry, setEntry }) {
  function setValue(value) {
    return setEntry(e => ({ ...e, ...value }));
  }

  return (
    <>
      <td>
        <input
          value={entry.name}
          onChange={event => setValue({ name: event.target.value })}
        />
      </td>
    </>
  );
}

EditEntry.propTypes = {
  entry: Envelope.isRequired,
  setEntry: PropTypes.func.isRequired,
};

function prepareInput(input) {
  return {
    name: input.name,
  };
}
