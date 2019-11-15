import React, { useState } from 'react';
import PropTypes from 'prop-types';

export default function ModalButton({ button, modal, openRef }) {
  const [show, setShow] = useState(false);
  const onHide = () => setShow(false);
  const onClick = () => setShow(true);
  const Button = button;
  const Modal = modal;
  if (openRef) {
    openRef.current = onClick;
  }
  return (
    <>
      <Button onClick={onClick} />
      <Modal onHide={onHide} show={show} />
    </>
  );
}

ModalButton.propTypes = {
  button: PropTypes.elementType.isRequired,
  modal: PropTypes.elementType.isRequired,
  openRef: PropTypes.shape({ current: PropTypes.any }),
};
