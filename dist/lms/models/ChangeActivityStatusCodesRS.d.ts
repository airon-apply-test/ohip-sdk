/**
 * OPERA Cloud Leisure Management API
 * APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ErrorType } from './ErrorType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Existing Operations Responses will eventually be modified to be extended from this type.
 * @export
 * @interface ChangeActivityStatusCodesRS
 */
export interface ChangeActivityStatusCodesRS {
    /**
     * Returning an empty element of this type indicates the successful processing of an message. This is used in conjunction with the Warning Type to report any warnings or business errors.
     * @type {object}
     * @memberof ChangeActivityStatusCodesRS
     */
    success?: object;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChangeActivityStatusCodesRS
     */
    warnings?: Array<WarningType>;
    /**
     * An error that occurred during the processing of a message.
     * @type {Array<ErrorType>}
     * @memberof ChangeActivityStatusCodesRS
     */
    errors?: Array<ErrorType>;
    /**
     * A reference for additional message identification, assigned by the requesting host system. When a request message includes an echo token the corresponding response message MUST include an echo token with an identical value.
     * @type {string}
     * @memberof ChangeActivityStatusCodesRS
     */
    echoToken?: string;
    /**
     * Indicates the creation date and time of the message in UTC using the following format specified by ISO 8601; YYYY-MM-DDThh:mm:ssZ with time values using the 24 hour clock (e.g. 20 November 2003, 1:59:38 pm UTC becomes 2003-11-20T13:59:38Z).
     * @type {Date}
     * @memberof ChangeActivityStatusCodesRS
     */
    timeStamp?: Date;
    /**
     * For all Opera versioned messages, the version of the message is indicated by a Opera Version value.
     * @type {string}
     * @memberof ChangeActivityStatusCodesRS
     */
    version?: string;
    /**
     * Allow end-to-end correlation of log messages with the corresponding Web service message throughout the processing of the Web service message.
     * @type {string}
     * @memberof ChangeActivityStatusCodesRS
     */
    correlationId?: string;
    /**
     * Indicates if the operation supports the ability to retry the request.
     * @type {boolean}
     * @memberof ChangeActivityStatusCodesRS
     */
    retryAllowed?: boolean;
    /**
     * Indicates if the operation supports the ability to force the retry request through OPERA services in the case where the external system continues to fail.
     * @type {boolean}
     * @memberof ChangeActivityStatusCodesRS
     */
    enforceAllowed?: boolean;
    /**
     * This attribute carries the user selected confirmation value on confirmation popup.
     * @type {boolean}
     * @memberof ChangeActivityStatusCodesRS
     */
    useLocalAllowed?: boolean;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ChangeActivityStatusCodesRS
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the ChangeActivityStatusCodesRS interface.
 */
export declare function instanceOfChangeActivityStatusCodesRS(value: object): boolean;
export declare function ChangeActivityStatusCodesRSFromJSON(json: any): ChangeActivityStatusCodesRS;
export declare function ChangeActivityStatusCodesRSFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeActivityStatusCodesRS;
export declare function ChangeActivityStatusCodesRSToJSON(value?: ChangeActivityStatusCodesRS | null): any;
