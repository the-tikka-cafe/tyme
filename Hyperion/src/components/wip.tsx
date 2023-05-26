import { AiOutlineLoading3Quarters } from 'react-icons/ai';
import { FaHammer } from 'react-icons/fa';

function WIP() {
  return (
    <div className="work-in-progress">
      <div className="illustration">
        <span className="loading-icon">
          <AiOutlineLoading3Quarters size={50} />
        </span>
        <span className="hammer-icon">
          <FaHammer />
        </span>
      </div>
      <div className="wip-text">This page is a work in progress.</div>
    </div>
  );
}

export default WIP;
