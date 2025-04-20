import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Git Simplifié',
    Svg: require('@site/static/img/undraw_programming_65t2.svg').default,
    description: (
      <>
        OwnGit te permet de manipuler une version allégée de Git pour mieux comprendre ses rouages internes, sans complexité inutile.
      </>
    ),
  },
  {
    title: 'Apprentissage Progressif',
    Svg: require('@site/static/img/undraw_educator_6dgp.svg').default,
    description: (
      <>
        Apprends les commandes essentielles étape par étape, dans un environnement fait pour l’expérimentation.
      </>
    ),
  },
  {
    title: 'Projet Fun & Ouvert',
    Svg: require('@site/static/img/undraw_fun-moments_x0p9.svg').default,
    description: (
      <>
        OwnGit est un projet open source fun, conçu pour ceux qui veulent apprendre en construisant. Forke-le, améliore-le et amuse-toi !
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
