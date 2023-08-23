/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DuplicateExternalSubscriptionsType } from './DuplicateExternalSubscriptionsType';
import type { InstanceLink } from './InstanceLink';
/**
 * Response object for fetching subscriptions where the same OPERA profile is linked to multiple external profiles within the same external system.
 * @export
 * @interface DuplicateExternalSubscriptions
 */
export interface DuplicateExternalSubscriptions {
    /**
     * Details of the OPERA Profile subscription to external system
     * @type {Array<DuplicateExternalSubscriptionsType>}
     * @memberof DuplicateExternalSubscriptions
     */
    duplicateExternalSubscriptionsList?: Array<DuplicateExternalSubscriptionsType>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof DuplicateExternalSubscriptions
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof DuplicateExternalSubscriptions
     */
    pageNumber?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof DuplicateExternalSubscriptions
     */
    maxFetchCount?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof DuplicateExternalSubscriptions
     */
    allRowsFetched?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof DuplicateExternalSubscriptions
     */
    totalRows?: number;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DuplicateExternalSubscriptions
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the DuplicateExternalSubscriptions interface.
 */
export declare function instanceOfDuplicateExternalSubscriptions(value: object): boolean;
export declare function DuplicateExternalSubscriptionsFromJSON(json: any): DuplicateExternalSubscriptions;
export declare function DuplicateExternalSubscriptionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): DuplicateExternalSubscriptions;
export declare function DuplicateExternalSubscriptionsToJSON(value?: DuplicateExternalSubscriptions | null): any;
