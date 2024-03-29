import { createUser, fetchUsers } from '@/redux/user';
import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import SweetAlert from '@/helpers/sweetalert';

export default function ModalUser() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirm_password, setConfirmPassword] = useState('');

  const dispatch = useDispatch();

  const handleChangeName = (e) => {
    setName(e.target.value);
  };

  const handleChangeEmail = (e) => {
    setEmail(e.target.value);
  };

  const handleChangePassword = (e) => {
    setPassword(e.target.value);
  };

  const handleConfirmPasswordChange = (e) => {
    setConfirmPassword(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (password !== confirm_password) {
      SweetAlert.error(
        'Error!!!',
        'Password and Confirm Password do not match'
      );
      return;
    }

    const formData = {
      name: name,
      email: email,
      password: password,
      confirm_password: confirm_password,
    };

    dispatch(createUser(formData))
      .then(() => {
        SweetAlert.success('Successfully', 'User created successfully!').then(
          () => dispatch(fetchUsers())
        );
      })
      .catch(() => {
        SweetAlert.error('Error!!', 'Failed to create user. Please try again.');
      });
  };

  return (
    <div
      className="modal fade text-left"
      id="user"
      tabIndex={-1}
      role="dialog"
      aria-labelledby="user"
      aria-hidden="true"
    >
      <div
        className="modal-dialog modal-dialog-centered modal-dialog-scrollable modal-lg"
        role="document"
      >
        <div className="modal-content">
          <div className="modal-header">
            <h4 className="modal-title" id="user">
              Add User
            </h4>
            <button
              type="button"
              className="close"
              data-bs-dismiss="modal"
              aria-label="Close"
            >
              <i data-feather="x" />
            </button>
          </div>
          <form onSubmit={handleSubmit}>
            <div className="modal-body">
              <label htmlFor="Name">Name: </label>
              <div className="form-group">
                <input
                  id="name"
                  type="text"
                  name="name"
                  placeholder="Name"
                  className="form-control"
                  value={name}
                  onChange={handleChangeName}
                />
              </div>

              <label htmlFor="email">Email: </label>
              <div className="form-group">
                <input
                  id="email"
                  type="text"
                  name="email"
                  placeholder="Email Address"
                  className="form-control"
                  value={email}
                  onChange={handleChangeEmail}
                />
              </div>
              <label htmlFor="password">Password: </label>
              <div className="form-group">
                <input
                  id="password"
                  type="password"
                  name="password"
                  placeholder="Password"
                  className="form-control"
                  value={password}
                  onChange={handleChangePassword}
                />
              </div>

              <label htmlFor="confirm_password">Confirm Password: </label>
              <div className="form-group">
                <input
                  id="confirm_password"
                  type="password"
                  name="password"
                  placeholder="Password"
                  className="form-control"
                  value={confirm_password}
                  onChange={handleConfirmPasswordChange}
                />
              </div>
            </div>
            <div className="modal-footer">
              <button
                type="button"
                className="btn btn-light-secondary"
                data-bs-dismiss="modal"
              >
                <i className="bx bx-x d-block d-sm-none" />
                <span className="d-none d-sm-block">Close</span>
              </button>
              <button
                type="submit"
                className="btn btn-primary ms-1"
                data-bs-dismiss="modal"
              >
                <i className="bx bx-check d-block d-sm-none" />
                <span className="d-none d-sm-block">Adding</span>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
