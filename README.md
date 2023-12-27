# Check Wehrner von Braun's circumlunar probe orbit

![circum lunar vehicle](lunar_probe_bw.png)

[Across the Space Frontier](https://space.nss.org/book-review-across-the-space-frontier/),
Joseph Kaplan, Wernher von Braun, Heinz Haber, Willey Ley, Oscar Schatchter, Fred Whipple,
edited by Cornelius Ryan, Viking Press, 1952.
part of a prose expostion,
and elaboration of von Braun's [The Mars Project](https://en.wikipedia.org/wiki/The_Mars_Project),
the technical appendix of [Project Mars](https://en.wikipedia.org/wiki/Project_Mars:_A_Technical_Tale),

Von Braun describes a lunar excursion,
to be undertaken as a way to ease into interplanetary travel,
and map the landing site for a lunar expedition.

## Vehicle Specs

![vehicle schematic](lunar_orbiter_35.png)

Line drawing from _Across the Space Frontier_

![von Braun sketch of vehicle](braun650.jpg)

Von Braun's own sketch, which includes a scale.

Unfortunately, von Braun did not see fit to numerically describe
the circumlunar vehicle as fully as he did the Three Stage Launch vehicle.
He does write:
"The propulsion unit is the same as that of the Third Stage".

The _PRINCIPAL DATA ON THE THREE_STAGE ROCKET SHIP_ table has these
relevant facts for the Third Stage:

* 220 ton thrust
* 0.77 ton/sec propellant consumption
* 9200 ft/sec exhaust velocity
* 286 sec I<sub>sp</sub>

In consistent SI units:

* Thrust 1.958E6 Kg m/s<sup>2</sup>
* Propellant consumption 700 kg/s
* Exhaust velocity 2805 m/s

There's a phrase in the book about leaving space station orbit requires "barely 2 minutes" of thrust.

From this, we can back out a total propellant consumption of 168,000 kg.
That's one 120 sec burn
to leave space station orbit, a very similar burn to circularize the eccentric
cislunar orbit.

Von Braun also gives a &#916;V requirement of 2835 m/s
to leave the space station's 1075-mile orbit,
getting into an orbit whose apogee is just outside the moon's orbit.

From these facts and the [Rocket Equation](),
we can back out a fully fueled mass of 193652 kg,
and a vehicle dry mass of 25652 kg.

## Orbit Diagram

![orbit diagram](circum_lunar_orbit.png)

## A Program for Verifying the Orbit

In accordance with [Gall's Law](http://principles-wiki.net/principles:gall_s_law),
I'm going to start small and work my way up to the hard tasks.
In the process, I intend to learn a little about numerical integration,
and refresh my 40-year-old knowledge of orbital mechanics.

Here's my project plan:

1. Numerically integrate a rocket thrusting for a period.
Double check that the change in velocity matches the Rocket Equation.
   * Wrote a [simulation](csm.go) based on the Apollo Command/Service module,
   since that's very well documented numerically.
2. Numerically integrate a circular orbit. That is,
using (vector) **F** = m **A**, where acceleration **A**
always points to the center of the earth, see if an object with
tangential velocity of 7069.5 meters/sec ends up in a 1075 mile circular orbit.
   * Wrote a [simulation](ss.go) of Wernher von Braun's polar-orbiting
   space station from _Across the Space Frontier_, using [symplectic Euler](https://en.wikipedia.org/wiki/Semi-implicit_Euler_method)
   numerical integration, 0.25 second time steps
     * orbit height 1730044.745 m (1075 mile), 7012.6 m/s velocity
     * Numerically integrates a nice, 7258 second circular orbit,
       orbital radius staying between 8100234 m and 8101993 m
   * Ran a long-period (230 days) [simulation](long_period1.go) of Wernher von Braun's polar-orbiting
     space station to see if numeric integration lost or gained energy due to
     numerical problems. It does not.
3. After one orbit, apply an impulsive velocity change,
determine if **F** = m **A** integration of earth's gravity
causes the vehicle to change to an elliptic orbit.
   * Wrote a simulation of von Braun's lunar excursion vehicle
   doing a 1000 meter/sec impulsive velocity change, again using symplectic Euler
   numerical integration.
     * perigee of 1730044.745 m (1075 mile), 8012.6 m/s velocity after the impulsive &#916;V
     * [vis viva](https://en.wikipedia.org/wiki/Vis-viva_equation) equation works out to an apogee of 1.523x10<sup>7</sup> m,
       numerically integrated orbit hits that exactly.
4. After one orbit, apply a non-impulsive velocity change,
determine if **F** = m **A** integration of earth's gravity
causes the vehicle to change to an elliptic orbit.
Use the same 1000 meter/sec &#916;V as in (3).
5. Try numeric integration of a [Hohmann transfer orbit]() works.
This requires two impulsive velocity changes.
6. See if a "barely 2 minute" continuous thrust and concomitant mass change
numerically integrates to a big ellipse that takes the vehicle
out to the radius of the Moon's orbit.
7. Try to put a 3rd body, the Moon, into the simulation of (5).

## Symplectic Euler Method

All of my numerical integrations work like this:

>// Start simulation where vehicle has some position (X,Y) relative
>// to center of the earth, and a vector velocity (V<sub>x</sub>,V<sub>y</sub>)
>// All motion constrained to plain of ecliptic, there are no Z components of anything.
>// I'm assuming a known, constant thrust and constant mass flow rate for the
>// rocket engines.
>for t := t0; t < t<sub>max</sub>; t += &#916;t {
>
>    // find distance from center of earth, which is at (0.0,0.0)
>    r = &radic;(X<sup>2</sup> + Y<sup>2</sup>)
>
>    // magnitude of attraction due to gravity
>    F<sub>grav</sub> = G M<sub>earth</sub>/(r<sup>2</sup>)
>
>    // X and Y direction components of gravitational forces
>    F<sub>x</sub> = (-X/r)F<sub>grav</sub>
>    F<sub>y</sub> = (-Y/r)F<sub>grav</sub>
>
>    // Vector components of thrust, assumed tangential to orbit
>    V = &radic;(V<sub>x</sub><sup>2</sup> + V<sub>y</sub><sup>2</sup>)
>    F<sub>x</sub> = Thrust V<sub>x</sub>/V
>    F<sub>y</sub> = Thrust V<sub>y</sub>/V
>
>    // Mass change for time step
>    M -=  &#916;M // constant mass flow rate, &#916;M does depend on time step
>
>    // Vector acceleration components
>    A<sub>x</sub = F<sub>x</sub>/M
>    A<sub>y</sub = F<sub>y</sub>/M
>
>    // Increment velocity components
>    V<sub>x</sub += A<sub>x</sub &#916;t
>    V<sub>y</sub += A<sub>y</sub &#916;t
>
>    // Increment position components
>    X += V<sub>x</sub &#916;t
>    Y += V<sub>y</sub &#916;t
>}

## References

* [Symplectic Euler method](https://www.mgaillard.fr/2021/07/11/euler-integration.html)
* [Earth and Moon facts](https://nssdc.gsfc.nasa.gov/planetary/factsheet/moonfact.html)
