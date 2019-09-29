import React from 'react';
import { List } from '../common/List/List';
import { useBudget } from '../context/budget/budget';
import { useQuery, useMutation } from '@apollo/react-hooks';
import {
  QUERY_CATEGORIES,
  UPDATE_CATEGORY,
  CREATE_CATEGORY,
} from './CategoriesList.gql';
import { Category } from '../../model/propTypes';
import PropTypes from 'prop-types';
import { QUERY_ENVELOPES } from '../EnvelopesList/EnvelopesList.gql';

export function CategoriesList() {
  const { id: budgetID } = useBudget();
  const { loading, error, data } = useQuery(QUERY_CATEGORIES, {
    variables: { budgetID },
  });
  const [updateCategory] = useMutation(UPDATE_CATEGORY, {
    refetchQueries: () => [
      {
        query: QUERY_CATEGORIES,
        variables: { budgetID },
      },
    ],
  });
  const [createCategory] = useMutation(CREATE_CATEGORY, {
    refetchQueries: () => [
      {
        query: QUERY_CATEGORIES,
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
        entries={data.categories}
        onCreate={input =>
          createCategory({
            variables: { budgetID, input: prepareInput(input) },
          })
        }
        onUpdate={(id, input) =>
          updateCategory({
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
      <th>Koperta</th>
    </>
  );
}

function ListEntry({ entry }) {
  return (
    <>
      <td>{entry.name}</td>
      <td>{entry.envelope.name}</td>
    </>
  );
}

ListEntry.propTypes = {
  entry: Category.isRequired,
};

function EditEntry({ entry, setEntry }) {
  const { id: budgetID } = useBudget();
  const { loading, error, data } = useQuery(QUERY_ENVELOPES, {
    variables: { budgetID },
  });

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

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
      <td>
        <select
          value={entry.envelopeID || (entry.envelope && entry.envelope.id)}
          onChange={event =>
            setValue({
              envelopeID: event.target.value,
            })
          }
        >
          <option></option>
          {data.envelopes.map(envelope => (
            <option key={envelope.id} value={envelope.id}>
              {envelope.name}
            </option>
          ))}
        </select>
      </td>
    </>
  );
}

EditEntry.propTypes = {
  entry: PropTypes.shape({
    name: PropTypes.string.isRequired,
    envelopeID: PropTypes.any,
    envelope: PropTypes.shape({
      id: PropTypes.any,
    }),
  }).isRequired,
  setEntry: PropTypes.func.isRequired,
};

function prepareInput(input) {
  return {
    name: input.name,
    envelopeID: input.envelopeID || input.envelope.id,
  };
}
