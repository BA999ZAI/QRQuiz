import React from "react";

const UserProfileModal = ({ onClose }) => {
  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>Профиль пользователя</h2>
        <div className="profile-info">
          <img
            src="https://via.placeholder.com/100"
            alt="Profile"
            className="profile-avatar"
          />
          <p>Email: ivan@example.com</p>
        </div>
        <button className="close-button" onClick={onClose}>
          Закрыть
        </button>
      </div>
    </div>
  );
};

export default UserProfileModal;