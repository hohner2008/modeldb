import Project from 'models/Project';
import * as React from 'react';
import { connect } from 'react-redux';
import ServiceFactory from '../../services/ServiceFactory';
import { fetchProjects, fetchSuccess } from '../../store/project';
import { IApplicationState, IConnectedReduxProps } from '../../store/store';
import ProjectWidget from '../ProjectWidget/ProjectWidget';
import styles from './Projects.module.css';

interface IPropsFromState {
  projects?: Project[] | null;
  loading: boolean;
}

type AllProps = IPropsFromState & IConnectedReduxProps;

class Projects extends React.Component<AllProps> {
  public componentDidMount() {
    const data = ServiceFactory.getDataService().getProjects();
    this.props.dispatch(fetchSuccess(data));
  }

  public render() {
    return (
      <div className={styles.projects}>
        <div className={styles.headPanel}>
          <div>Projects</div>
          <div />
          <div />
        </div>
        <div className={styles.widgets_list}>
          {this.props.projects ? this.props.projects.map((proj, i) => <ProjectWidget project={proj} key={i} />) : ''}
        </div>
      </div>
    );
  }
}

const mapStateToProps = ({ projects }: IApplicationState) => ({
  loading: projects.loading,
  projects: projects.data
});

export default connect<IPropsFromState, {}, {}, IApplicationState>(mapStateToProps)(Projects);
