import { Outlet, useLocation, useNavigate, useParams } from "@remix-run/react";
import { DateTime } from "luxon";
import { ArrowLeftIcon, ArrowRightIcon } from "@heroicons/react/24/solid";
import { useMemo } from "react";
import { LoaderFunction } from "@remix-run/node";
import { authenticator } from "~/services/auth.server";

export const loader: LoaderFunction = ({ request }) =>
  authenticator.authenticate("microsoft", request, { failureRedirect: "/" });

export default function EntryIndex() {
  const params = useParams();
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const date = useMemo(() => {
    if (!params.id) {
      return DateTime.now().toISODate().toFormat("dd LLL");
    }

    return DateTime.fromISO(params.id).toFormat("dd LLL");
  }, [params]);

  const handleDate = (dir: number) => {
    var date = params.id;
    if (!date) return;

    navigate(
      "/entry/" + DateTime.fromISO(date).plus({ days: dir }).toISODate()
    );
  };

  return (
    <>
      <div className="my-6">
        <div className="flex justify-around items-center">
          <ArrowLeftIcon
            className="w-8 h-8 cursor-pointer"
            onClick={() => handleDate(-1)}
          />
          <span className="font-semibold text-lg">{date}</span>
          <ArrowRightIcon
            className="w-8 h-8 cursor-pointer"
            onClick={() => handleDate(1)}
          />
        </div>
      </div>
      <div className="mx-4">
        <Outlet key={pathname} />
      </div>
    </>
  );
}
